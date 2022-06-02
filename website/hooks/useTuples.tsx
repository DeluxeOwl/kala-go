import create from "zustand";
import { defaultPc, defaultTuples } from "../ideconfigs/configs";
import PermissionCheck from "../types/permissionCheck";
import Subject from "../types/subject";
import Tuple from "../types/tuple";
import { showError } from "../util/notifications";

interface TuplesState {
  tuples: Tuple[];
  permissionChecks: PermissionCheck[];
  addPermissionCheck: (pc: PermissionCheck) => void;
  removePermissionCheck: (pc: PermissionCheck) => void;
  updatePermissionStatus: (
    pc: PermissionCheck,
    permission: boolean | undefined,
    logs?: string[]
  ) => void;
  addTuple: (tuple: Tuple) => void;
  removeTuple: (tuple: Tuple) => void;
  getUniqueSubjects: () => Subject[];
  setState: (tuples: Tuple[], pcs: PermissionCheck[]) => void;
}

const useTuples = create<TuplesState>((set, get) => ({
  tuples: defaultTuples,
  permissionChecks: defaultPc,

  setState: (tuples, pcs) =>
    set((state) => ({ ...state, tuples: tuples, permissionChecks: pcs })),

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
  updatePermissionStatus: (
    pc: PermissionCheck,
    permission: boolean | undefined,
    logs?: string[]
  ) => {
    set((state) => {
      return {
        permissionChecks: state.permissionChecks.map((p) =>
          p.subject.name === pc.subject.name &&
          p.subject.type === pc.subject.type &&
          p.permission === pc.permission &&
          p.resource.name === pc.resource.name &&
          p.resource.type === pc.resource.type
            ? { ...p, hasPermission: permission, logs: logs }
            : p
        ),
      };
    });
  },
  addPermissionCheck: (pc) =>
    set((state) => {
      const pcExists = state.permissionChecks.some(
        (p) =>
          p.subject.name === pc.subject.name &&
          p.subject.type === pc.subject.type &&
          p.permission === pc.permission &&
          p.resource.name === pc.resource.name &&
          p.resource.type === pc.resource.type
      );
      if (pcExists) {
        showError("Permission check already exists");
        return { permissionChecks: state.permissionChecks };
      }

      return { permissionChecks: [pc, ...state.permissionChecks] };
    }),
  removePermissionCheck: (pc) =>
    set((state) => ({
      permissionChecks: state.permissionChecks.filter(
        (p) =>
          !(
            p.subject.name === pc.subject.name &&
            p.subject.type === pc.subject.type &&
            p.permission === pc.permission &&
            p.resource.name === pc.resource.name &&
            p.resource.type === pc.resource.type
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
