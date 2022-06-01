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
    labelBgStyle: { fill: "#1acc92", color: "#fff", fillOpacity: 0.7 },
    markerEnd: {
      type: MarkerType.ArrowClosed,
      color: "#1acc92",
      height: 30,
      width: 30,
    },
    target: targetId,
    style: {
      stroke: "#1acc92",
      opacity: 0.5,
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
    labelBgStyle: { fill: "#ffa887", color: "#fff", fillOpacity: 0.7 },
    markerEnd: {
      type: MarkerType.ArrowClosed,
      color: "#ffa887",
      height: 30,
      width: 30,
    },
    target: targetId,
    style: {
      stroke: "#ffa887",
      opacity: 0.5,
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
    style: {
      strokeWidth: 4,
    },
    markerEnd: {
      type: MarkerType.Arrow,
    },
    animated: true,
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

      height: 30,
      width: 30,
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

      height: 30,
      width: 30,
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

    markerEnd: {
      type: MarkerType.ArrowClosed,
    },
    target: targetId,
    style: {
      stroke: "#008000",
      strokeWidth: 4,
    },
  };
};
const relToTcEdge = (id: string, sourceId: string, targetId: string): Edge => {
  return {
    id: id,
    source: sourceId,
    style: {
      strokeWidth: 4,
    },
    markerEnd: {
      type: MarkerType.Arrow,
    },
    target: targetId,
    animated: true,
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
    markerEnd: {
      type: MarkerType.ArrowClosed,
      color: "#1f77b4",
    },
    style: {
      stroke: "#1f77b4",
      strokeWidth: 4,
    },
    animated: true,
  };
};

const binaryExprOperatorEdge = (
  id: string,
  sourceId: string,
  targetId: string,
  operator: string
): Edge => {
  let colorDependingOnOperator: string;
  switch (operator) {
    case "|":
      colorDependingOnOperator = "#008000";
      break;

    case "&":
      colorDependingOnOperator = "#f29d02";
      break;

    case "!":
      colorDependingOnOperator = "#A7171A";
      break;

    default:
      colorDependingOnOperator = "#fff";
      break;
  }

  return {
    id: id,
    source: sourceId,
    target: targetId,
    markerEnd: {
      type: MarkerType.ArrowClosed,
      color: colorDependingOnOperator,
    },
    style: {
      stroke: colorDependingOnOperator,
      strokeWidth: 4,
    },
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
