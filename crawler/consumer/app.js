const dotenv = require('dotenv')
const amqp = require('amqplib');
const neo4j = require('neo4j-driver')


dotenv.config({
  path: '../../.env'
})

const rHost = process.env.RABBITMQ_HOST
const rVHost = process.env.RABBITMQ_VHOST
const rUser = process.env.RABBITMQ_USER
const rPwd = process.env.RABBITMQ_PWD

const nHost = process.env.NEO4J_HOST
const nUser = process.env.NEO4J_USER
const nPwd = process.env.NEO4J_PWD

const amqpURL = `amqp://${rUser}:${rPwd}@${rHost}${rVHost}`
const driver = neo4j.driver(nHost, neo4j.auth.basic(nUser, nPwd))


async function consume() {
  const conn = await amqp.connect(amqpURL)
  const ch = await conn.createChannel()
  const q = 'music-creation'
  await conn.createChannel()
  await ch.assertQueue(q, {durable: true});
  await ch.consume(q, async (msg) => {
    const session = driver.session()
    const data = JSON.parse(msg.content.toString());

    console.log(`[consume] received ${JSON.stringify(data)}`)

    if (!data.releases)  {
      ch.ack(msg)
      return
    }

    for (const release of data.releases) {
      await session.run(
        `MERGE(b:Band{name: $band})
        MERGE(a:Album{name: $album})
        MERGE(b)-[:RELEASE{released_at: $date}]->(a)`,
        { band: data.artist.name, album: release.title, date: release.date }
      )
      
      if (!release.media) {
        ch.ack(msg)
        return
      }

      for (const media of release.media) {
        if (!media.tracks) {
          ch.ack(msg)
          return
        }
        for (const track of media.tracks) {
          await session.run(
            `MATCH(a:Album{name:$album})
            MERGE(m:Music{name: $name})
            MERGE(m)<-[:HAS_MUSIC{order: $order}]-(a)`,
            { name: track.title, order: track.position, album: release.title }
          )
        }
      }
    }

    ch.ack(msg)
  })
}

consume()