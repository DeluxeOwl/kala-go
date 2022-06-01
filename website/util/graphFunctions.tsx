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
    type: "typeConfigNode",
    data: { label: label },
    position: position,
  };
};

const relNode = (id: string, label: string, position: Point): Node => {
  return {
    id: id,
    type: "relationNode",
    data: { label: label },
    position: position,
  };
};

const permNode = (id: string, label: string, position: Point): Node => {
  return {
    id: id,
    type: "permissionNode",
    data: { label: label },
    position: position,
  };
};

const permEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: id,
    source: sourceId,
    label: "permission",
    labelBgPadding: [8, 4],
    labelBgBorderRadius: 4,
    labelBgStyle: { fill: "#7F00FF ", color: "#fff", fillOpacity: 0.7 },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: targetId,
    style: {
      stroke: "violet",
    },
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
    type: "relComposedNode",
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
const permNotEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: `${id}/not`,
    source: sourceId,
    style: {
      stroke: "red",
    },
    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: `${targetId}`,
  };
};
const permDirectEdge = (
  id: string,
  sourceId: string,
  targetId: string
): Edge => {
  return {
    id: `${id}/direct`,
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
    target: `${targetId}/not`,
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

const memberExprEdge = (
  id: string,
  sourceId: string,
  targetId: string
): Edge => {
  return {
    id: id,
    source: sourceId,
    target: targetId,
    style: {
      stroke: "cyan",
    },
    animated: true,
  };
};

const binaryExprOperatorEdge = (
  id: string,
  sourceId: string,
  targetId: string
): Edge => {
  return {
    id: id,
    source: sourceId,
    target: targetId,
    style: {
      stroke: "cyan",
    },
    animated: true,
  };
};

// change depending on label !, &, |
const binaryExprOperatorNode = (
  id: string,
  label: string,
  position: Point
): Node => {
  return {
    id: id,
    type: "binaryExprOperatorNode",
    data: { label: label },
    position: position,
  };
};

function randomIntFromInterval(min: number, max: number) {
  // min and max included
  return Math.floor(Math.random() * (max - min + 1) + min);
}
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
  binaryExprOperatorEdge,
  permNode,
  permEdge,
  permDirectEdge,
  permNotEdge,
  randomIntFromInterval,
  memberExprEdge,
  binaryExprOperatorNode,
};
