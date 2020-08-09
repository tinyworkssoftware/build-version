use build_version_db;
INSERT INTO tbl_organisation(id, name, plan_type)
VALUES(uuid(), 'test-org', (SELECT id from tbl_plan_type WHERE tbl_plan_type.name = 'free'));

INSERT into tbl_project(id, name, organisation, access_code)
VALUES(uuid(), 'test-project', (SELECT id from tbl_organisation WHERE name = 'test-org'), uuid());

select * from tbl_project;