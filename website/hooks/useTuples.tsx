import create from "zustand";
import Subject from "../types/subject";
import Tuple from "../types/tuple";
import { showError } from "../util/notifications";

const initialTuples: Tuple[] = [
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

interface TuplesState {
  tuples: Tuple[];
  addTuple: (tuple: Tuple) => void;
  removeTuple: (tuple: Tuple) => void;
  getUniqueSubjects: () => Subject[];
}

const useTuples = create<TuplesState>((set, get) => ({
  tuples: initialTuples,
  addTuple: (tuple) =>
    set((state) => {
      const tupleExists = state.tuples.some(
        (t) =>
          t.subject.name === tuple.subject.name &&
          t.subject.type === tuple.subject.type &&
          t.relation === tuple.relation &&
          t.resource.name === tuple.resource.name &&
          t.resource.type === tuple.resource.type
      );
      if (tupleExists) {
        showError("Tuple already exists");
        return { tuples: state.tuples };
      }

      return { tuples: [tuple, ...state.tuples] };
    }),
  removeTuple: (tuple) =>
    set((state) => ({
      tuples: state.tuples.filter(
        (t) =>
          !(
            t.subject.name === tuple.subject.name &&
            t.subject.type === tuple.subject.type &&
            t.relation === tuple.relation &&
            t.resource.name === tuple.resource.name &&
            t.resource.type === tuple.resource.type
          )
      ),
    })),
  getUniqueSubjects: () => {
    const subjects: Subject[] = [];
    const flag: any = {};

    get().tuples.forEach((t: Tuple) => {
      let subjectType = t.subject.type;
      let subjectName = t.subject.name;
      let resourceType = t.resource.type;
      let resourceName = t.resource.name;

      if (subjectName.includes("#")) {
        subjectName = subjectName.split("#")[0];
      }
      if (resourceName.includes("#")) {
        resourceName = resourceName.split("#")[0];
      }

      if (!flag[subjectType + ":" + subjectName]) {
        subjects.push({
          type: subjectType,
          name: subjectName,
        });
        flag[subjectType + ":" + subjectName] = true;
      }
      if (!flag[resourceType + ":" + resourceName]) {
        subjects.push({
          type: resourceType,
          name: resourceName,
        });
        flag[resourceType + ":" + resourceName] = true;
      }
    });
    return subjects;
  },
}));

export default useTuples;
