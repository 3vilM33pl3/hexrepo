PRAGMA foreign_keys = ON;
CREATE TABLE hexrepo(id TEXT UNIQUE);
CREATE TABLE hexdata(hexid TEXT, key TEXT, value TEXT, FOREIGN KEY(hexid) REFERENCES hexrepo(id));
CREATE TABLE hexmap(id INTEGER, x INTEGER, y INTEGER, z INTEGER, hexid TEXT, FOREIGN KEY(hexid) REFERENCES hexrepo(id));
CREATE UNIQUE INDEX hmi ON hexmap(id);
CREATE TABLE mapdata(id INTEGER, key TEXT, value TEXT, FOREIGN KEY(id) REFERENCES hexmap(id));
CREATE UNIQUE INDEX mdi ON mapdata(id);