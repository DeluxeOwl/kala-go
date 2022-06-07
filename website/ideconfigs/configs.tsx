import PermissionCheck from "../types/permissionCheck";
import Tuple from "../types/tuple";
const defaultYaml = `
type: user
---
type: group
relations:
  member: user

---
type: folder
relations:
  reader: user | group#member

---
type: document
relations:
  parent_folder: folder
  writer: user
  reader: user
permissions:
  read: reader | writer | parent_folder.reader
  read_and_write: reader & writer
  read_only: reader & !writer
`;
const defaultTuples: Tuple[] = [
  {
    subject: {
      type: "user",
      name: "anna",
    },
    relation: "reader",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "anna",
    },
    relation: "writer",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "folder",
      name: "secret_folder",
    },
    relation: "parent_folder",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "dev",
    },
  },
  {
    subject: {
      type: "group",
      name: "dev#member",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "group",
      name: "test_group#member",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "dev",
    },
  },
];

const defaultPc: PermissionCheck[] = [
  {
    subject: {
      type: "user",
      name: "john",
    },
    permission: "read",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "anna",
    },
    permission: "read",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    permission: "read",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "anna",
    },
    permission: "read_only",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "anna",
    },
    permission: "read_and_write",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    permission: "read_and_write",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    permission: "read_and_write",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
];

const rbacYaml = `
type: user
---

type: document
relations:
  writer: user
  reader: user
permissions:
  edit: writer
  view: reader | writer
`;

const rbacTuples: Tuple[] = [
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "writer",
    resource: {
      type: "document",
      name: "some_doc",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "reader",
    resource: {
      type: "document",
      name: "some_doc",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    relation: "reader",
    resource: {
      type: "document",
      name: "some_another_doc",
    },
  },
];

const rbacPc: PermissionCheck[] = [
  {
    subject: {
      type: "user",
      name: "john",
    },
    permission: "view",
    resource: {
      type: "document",
      name: "some_doc",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    permission: "view",
    resource: {
      type: "document",
      name: "some_doc",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    permission: "view",
    resource: {
      type: "document",
      name: "some_another_doc",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    permission: "view",
    resource: {
      type: "document",
      name: "some_another_doc",
    },
  },
];

const gdriveYaml = `
type: user
---
type: group
relations:
  member: user

---
type: folder
relations:
  owner: user
  parent: folder
  viewer: user | group#member
permissions:
  create_file: owner
  view: viewer | owner | parent.viewer

---
type: doc
relations:
  owner: user
  parent: folder
  viewer: user | group#member
permissions:
  change_owner: owner | parent.owner
  read: viewer | owner | parent.viewer
  share: owner
  write: owner | parent.owner
`;

const gdriveTuples: Tuple[] = [
  {
    subject: {
      type: "folder",
      name: "produs-2022",
    },
    relation: "parent",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "maria",
    },
    relation: "viewer",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "anne",
    },
    relation: "owner",
    resource: {
      type: "folder",
      name: "produs-2022",
    },
  },
  {
    subject: {
      type: "group",
      name: "acme_group#member",
    },
    relation: "viewer",
    resource: {
      type: "folder",
      name: "produs-2022",
    },
  },
  {
    subject: {
      type: "user",
      name: "anne",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "contoso",
    },
  },
  {
    subject: {
      type: "user",
      name: "darius",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "acme_group",
    },
  },
];

const gdrivePc: PermissionCheck[] = [
  {
    subject: {
      type: "user",
      name: "anne",
    },
    permission: "write",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "anne",
    },
    permission: "change_owner",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "maria",
    },
    permission: "change_owner",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "darius",
    },
    permission: "read",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "darius",
    },
    permission: "write",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
  {
    subject: {
      type: "user",
      name: "daniel",
    },
    permission: "read",
    resource: {
      type: "doc",
      name: "finalizare_studii.docx",
    },
  },
];

export {
  defaultTuples,
  defaultPc,
  defaultYaml,
  rbacPc,
  rbacYaml,
  rbacTuples,
  gdriveYaml,
  gdriveTuples,
  gdrivePc,
};
