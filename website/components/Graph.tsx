// @ts-nocheck

import { useEffect } from "react";
import ReactFlow, {
  Background,
  Controls,
  Edge,
  Node,
  useEdgesState,
  useNodesState,
} from "react-flow-renderer";
import {
  cosDegrees,
  relComposedEdge,
  relComposedNode,
  relComposedSubrelEdge,
  relEdge,
  relNode,
  relToTcEdge,
  sinDegrees,
  tcNode,
} from "../util/graphFunctions";

type Point = {
  x: number;
  y: number;
};

type NodesAndEdges = {
  nodes: Node[];
  edges: Edge[];
};

const refValueDelim = " | ";
const refSubrelationDelim = "#";
const parentRelDelim = ".";

// TODO: calculate some stuff here to look good
const getNodes = (graph: any): Node[] => {
  let nodes: Node[] = [];
  let edges: Edge[] = [];

  // https://stackoverflow.com/questions/5300938/calculating-the-position-of-points-in-a-circle
  let radius = 500;
  let degrees: number = 360;
  if (graph.length > 0) {
    degrees = 360 / graph.length;
  }

  graph.forEach((tc, i) => {
    const tcId = `tc/${tc.name}`;
    const tcLabel = tc.name;

    const computedDgTc = degrees * (i + 1);

    const tcPosition: Point = {
      x: radius * cosDegrees(computedDgTc),
      y: radius * sinDegrees(computedDgTc),
    };

    // tcPoint.x += 250;

    nodes.push(tcNode(tcId, tcLabel, tcPosition));

    const tcEdges = tc?.edges;

    for (const prop in tcEdges) {
      if (prop === "relations") {
        const relations = tcEdges[prop];

        const computedDgRel = degrees * (i + 1);

        let relPoint: Point = {
          x: tcPosition.x + radius * cosDegrees(computedDgRel),
          y: tcPosition.y + radius * sinDegrees(computedDgRel),
        };

        relations.forEach((rel, i) => {
          const relId = `${tcId}/rel/${rel.name}`;
          const edgeId = `${tcId}-${relId}`;
          const relLabel = rel.name;
          const relPosition: Point = {
            x: relPoint.x,
            y: relPoint.y,
          };

          relPoint.y += 250;
          relPoint.x += 250;

          edges.push(relEdge(edgeId, tcId, relId));

          nodes.push(relNode(relId, relLabel, relPosition));

          // Composed relation
          const relValue: string = rel?.value;
          if (relValue.includes(refValueDelim)) {
            nodes.push(
              relComposedNode(`${relId}`, {
                x: relPosition.x + 250,
                y: relPosition.y + 250,
              })
            );

            edges.push(relComposedEdge(edgeId, relId, relId));

            for (const [i, composedRel] of relValue
              .split(refValueDelim)
              .entries()) {
              if (composedRel.includes(refSubrelationDelim)) {
                const s: string = composedRel.split(refSubrelationDelim);
                edges.push(
                  relComposedSubrelEdge(
                    `${edgeId}/or/${composedRel}/${i}`,
                    `${relId}/or`,
                    `tc/${s[0]}/rel/${s[1]}`
                  )
                );
              } else {
                edges.push(
                  relComposedSubrelEdge(
                    `${edgeId}/or/${composedRel}/${i}`,
                    `${relId}/or`,
                    `tc/${composedRel}`
                  )
                );
              }
            }
          } else {
            edges.push(
              relToTcEdge(
                `${tcId}-${relId}-${relValue}`,
                relId,
                `tc/${relValue}`
              )
            );
          }
        });
      }
      if (prop === "permissions") {
        const permissions = tcEdges[prop];
      }
      if (prop === "subjects") {
        const subjects = tcEdges[prop];
      }
    }
  });

  const ne: NodesAndEdges = {
    nodes: nodes,
    edges: edges,
  };

  return ne;
};

type GraphProps = {
  data: any;
};

const Graph = ({ data }: GraphProps) => {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  useEffect(() => {
    const parsedData = getNodes(data);
    setNodes(parsedData.nodes);
    setEdges(parsedData.edges);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  return (
    <ReactFlow
      nodes={nodes}
      onNodesChange={onNodesChange}
      edges={edges}
      onEdgesChange={onEdgesChange}
      nodesConnectable={false}
      connectionMode={"loose"}
      fitView
    >
      <Controls />
      <Background />
    </ReactFlow>
  );
};

export default Graph;
