import { Edge, MarkerType, Node } from "react-flow-renderer";

type Point = {
  x: number;
  y: number;
};

function sinDegrees(angleDegrees: number) {
  return Math.sin((angleDegrees * Math.PI) / 180);
}
function cosDegrees(angleDegrees: number) {
  return Math.cos((angleDegrees * Math.PI) / 180);
}

const tcNode = (id: string, label: string, position: Point): Node => {
  return {
    id: id,
    data: { label: label },
    position: position,
  };
};

const relNode = (id: string, label: string, position: Point): Node => {
  return {
    id: id,
    data: { label: label },
    position: position,
  };
};

const relEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: id,
    source: sourceId,
    label: "relation",
    labelBgPadding: [8, 4],
    labelBgBorderRadius: 4,
    labelBgStyle: { fill: "#FFCC00", color: "#fff", fillOpacity: 0.7 },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: targetId,
    style: {
      stroke: "yellow",
    },
  };
};

const relComposedNode = (id: string, position: Point): Node => {
  return {
    id: `${id}/or`,
    data: { label: "|" },
    position: position,
  };
};
const relComposedEdge = (
  id: string,
  sourceId: string,
  targetId: string
): Edge => {
  return {
    id: `${id}/or`,
    source: sourceId,
    label: "includes",
    labelBgPadding: [8, 4],
    labelBgStyle: {
      fill: "green",
      color: "#fff",
      fillOpacity: 0.7,
    },
    style: {
      stroke: "green",
    },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: `${targetId}/or`,
  };
};
const relComposedSubrelEdge = (
  id: string,
  sourceId: string,
  targetId: string
): Edge => {
  return {
    id: id,
    source: sourceId,
    label: "OR",
    labelBgPadding: [8, 4],
    labelBgBorderRadius: 4,
    labelBgStyle: {
      fill: "#0000FF",
      color: "#fff",
      fillOpacity: 0.7,
    },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: targetId,
    style: {
      stroke: "blue",
    },
  };
};
const relToTcEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: id,
    source: sourceId,
    label: "includes",
    labelBgPadding: [8, 4],
    labelBgBorderRadius: 4,
    labelBgStyle: {
      fill: "green",
      color: "#fff",
      fillOpacity: 0.7,
    },
    style: {
      stroke: "green",
    },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: targetId,
  };
};

const subjectNode = (id: string, label: string, position: Point): Node => {
  return {
    id: id,
    data: { label: label },
    position: position,
  };
};

const subjectEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: id,
    source: sourceId,
    target: targetId,
    style: {
      stroke: "red",
    },
  };
};

export {
  sinDegrees,
  cosDegrees,
  relComposedEdge,
  relComposedNode,
  relComposedSubrelEdge,
  relToTcEdge,
  tcNode,
  relNode,
  relEdge,
  subjectNode,
  subjectEdge,
};
