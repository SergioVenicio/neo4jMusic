MERGE (guitar:Instrument{name: "Guitar"})
MERGE (batery:Instrument{name: "Batery"})
MERGE (bassGuitar:Instrument{name: "Bass Guitar"})
MERGE (vocal:Instrument{name: "Vocal"})

MERGE (lemmy:Musician{name: "Lemmy Kilmister"})
MERGE (lemmy)-[:Play]-(bassGuitar)
MERGE (lemmy)-[:Play]-(vocal)

MERGE (phil:Musician{name: "Phil Campbell"})
MERGE (phil)-[:Play]-(guitar)

MERGE (mikkey:Musician{name: "Mikkey Dee"})
MERGE (mikkey)-[:Play]-(batery)

MERGE (motorhead:Band{name: "Motörhead"})
MERGE (lemmy)-[:Belongs{start_at: 1975, end_at: 2015}]-(motorhead)
MERGE (phil)-[:Belongs{start_at: 1984, end_at: 2015}]-(motorhead)
MERGE (mikkey)-[:Belongs{start_at: 1992, end_at: 2015}]-(motorhead)

MERGE (aceOfSpades:Album{name: "Ace of Spades", released_at: 1980})
MERGE (motorhead)-[:Release]-(aceOfSpades)


MERGE (music1:Music{name: "Ace of Spades", duration: 149.4, order: 1})
MERGE (music2:Music{name: "Love Me Like a Reptile", duration: 193.8, order: 2})
MERGE (music3:Music{name: "Shoot You in the Back", duration: 137.4, order: 3})
MERGE (music4:Music{name: "Live to Win", duration: 202.2, order: 4})
MERGE (music5:Music{name: "Fast and Loose", duration: 193.8, order: 5})
MERGE (music6:Music{name: "(We Are) The Road Crew", duration: 187.2, order: 6})
MERGE (music7:Music{name: "Fire Fire", duration: 146.4, order: 7})
MERGE (music8:Music{name: "Jailbait", duration: 199.8, order: 8})
MERGE (music9:Music{name: "Dance", duration: 142.8, order: 9})
MERGE (music10:Music{name: "Bite the Bullet", duration: 82.8, order: 10})
MERGE (music11:Music{name: "The Chase Is Better Than the Catch", duration: 250.8, order: 11})
MERGE (music12:Music{name: "The Hammer", duration: 148.8, order: 12})

MERGE (music1)-[:Belongs]-(aceOfSpades)
MERGE (music2)-[:Belongs]-(aceOfSpades)
MERGE (music3)-[:Belongs]-(aceOfSpades)
MERGE (music4)-[:Belongs]-(aceOfSpades)
MERGE (music5)-[:Belongs]-(aceOfSpades)
MERGE (music6)-[:Belongs]-(aceOfSpades)
MERGE (music7)-[:Belongs]-(aceOfSpades)
MERGE (music8)-[:Belongs]-(aceOfSpades)
MERGE (music9)-[:Belongs]-(aceOfSpades)
MERGE (music10)-[:Belongs]-(aceOfSpades)
MERGE (music11)-[:Belongs]-(aceOfSpades)
MERGE (music12)-[:Belongs]-(aceOfSpades)


MATCH(i:Instrument)<-[:Play]-(p:Musician)-[:Belongs]->(b:Band{name: "Motörhead"})-[:Release]->(a:Album)<-[:Belongs]-(m:Music)
RETURN *

MATCH(i:Instrument)<-[:Play]-(p:Musician)-[:Belongs]->(b:Band{name: "Motörhead"})-[:Release]->(a:Album)<-[:Belongs]-(m:Music)
RETURN SUM(m.duration) AS  durationInSeconds, a.name AS Album, Collect(DISTINCT m.name) AS Musics, Collect(DISTINCT p.name) AS Musicians