create table casbin_rule(p_type varchar, v0 varchar, v1 varchar, v2 varchar, v3 varchar, v4 varchar, v5 varchar);

INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/swagger/*', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/login', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/posts', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/post/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/liga/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/ligas', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/game/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'unauthorized', '/v1/games', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/posts', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/admin/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/ligas', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/liga/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'PUT');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/post', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/games', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/game/{id}', 'DELETE');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/club', 'POST');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/club/{id}', 'GET');
INSERT INTO casbin_rule(p_type, v0, v1, v2) VALUES('P', 'authorized', '/v1/clubs', 'GET');

