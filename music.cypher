MERGE (guitar:Instrument{name: "Guitar"})
MERGE (batery:Instrument{name: "Batery"})
MERGE (bassGuitar:Instrument{name: "Bass Guitar"})
MERGE (vocal:Instrument{name: "Vocal"})

MERGE (lemmy:Musician{name: "Lemmy Kilmister"})
MERGE (lemmy)-[:PLAYS]->(bassGuitar)
MERGE (lemmy)-[:PLAYS]->(vocal)

MERGE (phil:Musician{name: "Phil Campbell"})
MERGE (phil)-[:PLAYS]->(guitar)

MERGE (mikkey:Musician{name: "Mikkey Dee"})
MERGE (mikkey)-[:PLAYS]->(batery)

MERGE (motorhead:Band{name: "Motörhead"})
MERGE (lemmy)-[:BELONGS{start_at: 1975, end_at: 2015}]->(motorhead)
MERGE (phil)-[:BELONGS{start_at: 1984, end_at: 2015}]->(motorhead)
MERGE (mikkey)-[:BELONGS{start_at: 1992, end_at: 2015}]->(motorhead)

MERGE (aceOfSpades:Album{name: "Ace of Spades"})
MERGE (motorhead)-[:RELEASE{released_at: 1980}]->(aceOfSpades)

MERGE (lemmy)-[:PLAYED{instruments: ["Vocal", "Bass Guitar"]}]->(aceOfSpades)
MERGE (phil)-[:PLAYED{instruments: ["Guitar"]}]->(aceOfSpades)
MERGE (mikkey)-[:PLAYED{instruments: ["Batery"]}]->(aceOfSpades)


MERGE (aceOfSpadesmusic1:Music{name: "Ace of Spades", duration: 149.4})
MERGE (aceOfSpadesmusic2:Music{name: "Love Me Like a Reptile", duration: 193.8})
MERGE (aceOfSpadesmusic3:Music{name: "Shoot You in the Back", duration: 137.4})
MERGE (aceOfSpadesmusic4:Music{name: "Live to Win", duration: 202.2})
MERGE (aceOfSpadesmusic5:Music{name: "Fast and Loose", duration: 193.8})
MERGE (aceOfSpadesmusic6:Music{name: "(We Are) The Road Crew", duration: 187.2})
MERGE (aceOfSpadesmusic7:Music{name: "Fire Fire", duration: 146.4})
MERGE (aceOfSpadesmusic8:Music{name: "Jailbait", duration: 199.8})
MERGE (aceOfSpadesmusic9:Music{name: "Dance", duration: 142.8})
MERGE (aceOfSpadesmusic10:Music{name: "Bite the Bullet", duration: 82.8})
MERGE (aceOfSpadesmusic11:Music{name: "The Chase Is Better Than the Catch", duration: 250.8})
MERGE (aceOfSpadesmusic12:Music{name: "The Hammer", duration: 148.8})

MERGE (aceOfSpadesmusic1)<-[:HAS_MUSIC{order: 1}]-(aceOfSpades)
MERGE (aceOfSpadesmusic2)<-[:HAS_MUSIC{order: 2}]-(aceOfSpades)
MERGE (aceOfSpadesmusic3)<-[:HAS_MUSIC{order: 3}]-(aceOfSpades)
MERGE (aceOfSpadesmusic4)<-[:HAS_MUSIC{order: 4}]-(aceOfSpades)
MERGE (aceOfSpadesmusic5)<-[:HAS_MUSIC{order: 5}]-(aceOfSpades)
MERGE (aceOfSpadesmusic6)<-[:HAS_MUSIC{order: 6}]-(aceOfSpades)
MERGE (aceOfSpadesmusic7)<-[:HAS_MUSIC{order: 7}]-(aceOfSpades)
MERGE (aceOfSpadesmusic8)<-[:HAS_MUSIC{order: 8}]-(aceOfSpades)
MERGE (aceOfSpadesmusic9)<-[:HAS_MUSIC{order: 9}]-(aceOfSpades)
MERGE (aceOfSpadesmusic10)<-[:HAS_MUSIC{order: 10}]-(aceOfSpades)
MERGE (aceOfSpadesmusic11)<-[:HAS_MUSIC{order: 11}]-(aceOfSpades)
MERGE (aceOfSpadesmusic12)<-[:HAS_MUSIC{order: 12}]-(aceOfSpades)


MERGE (motorheadAlbum:Album{name: "Motörhead"})
MERGE (motorhead)-[:RELEASE{released_at: 1977}]->(motorheadAlbum)

MERGE (lemmy)-[:PLAYED{instruments: ["Vocal", "Bass Guitar"]}]->(motorheadAlbum)
MERGE (phil)-[:PLAYED{instruments: ["Guitar"]}]->(motorheadAlbum)
MERGE (mikkey)-[:PLAYED{instruments: ["Batery"]}]->(motorheadAlbum)

MERGE (music1:Music{name: "Motörhead", duration: 186.6})
MERGE (music2:Music{name: "Vibrator", duration: 201.6})
MERGE (music3:Music{name: "Lost Johnny", duration: 248.4})
MERGE (music4:Music{name: "Iron horse / Born To Lose", duration: 312})
MERGE (music5:Music{name: "White Line Fever", duration: 142.2})
MERGE (music6:Music{name: "Keep Us On The Road", duration: 333})
MERGE (music7:Music{name: "The Watcher", duration: 256.2})
MERGE (music8:Music{name: "Train Kept A-Rollin'", duration: 190.2})

MERGE (music1)<-[:HAS_MUSIC{order: 1}]-(motorheadAlbum)
MERGE (music2)<-[:HAS_MUSIC{order: 2}]-(motorheadAlbum)
MERGE (music3)<-[:HAS_MUSIC{order: 3}]-(motorheadAlbum)
MERGE (music4)<-[:HAS_MUSIC{order: 4}]-(motorheadAlbum)
MERGE (music5)<-[:HAS_MUSIC{order: 5}]-(motorheadAlbum)
MERGE (music6)<-[:HAS_MUSIC{order: 6}]-(motorheadAlbum)
MERGE (music7)<-[:HAS_MUSIC{order: 7}]-(motorheadAlbum)
MERGE (music8)<-[:HAS_MUSIC{order: 8}]-(motorheadAlbum)



MATCH(i:Instrument)<-[:Play]-(p:Musician)-[:Belongs]->(b:Band{name: "Motörhead"})-[:Release]->(a:Album)<-[:Belongs]-(m:Music)
RETURN *

MATCH(i:Instrument)<-[:Play]-(p:Musician)-[:Belongs]->(b:Band{name: "Motörhead"})-[:Release]->(a:Album)<-[:Belongs]-(m:Music)
RETURN SUM(m.duration) AS  durationInSeconds, a.name AS Album, Collect(DISTINCT m.name) AS Musics, Collect(DISTINCT p.name) AS Musicians