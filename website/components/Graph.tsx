// @ts-nocheck

import { useEffect } from "react";
import ReactFlow, {
  Background,
  Controls,
  Edge,
  MarkerType,
  Node,
  useEdgesState,
  useNodesState,
} from "react-flow-renderer";

type Point = {
  x: number;
  y: number;
};

type NodesAndEdges = {
  nodes: Node[];
  edges: Edge[];
};

function randomIntFromInterval(min, max) {
  // min and max included
  return Math.floor(Math.random() * (max - min + 1) + min);
}

const getNodes = (data: any): Node[] => {
  let nodes: Node[] = [];
  let edges: Edge[] = [];

  let tcPoint: Point = {
    x: 0,
    y: 0,
  };

  data.forEach((tc, i) => {
    const tcId = `tc/${tc.name}`;
    const tcLabel = tc.name;
    const tcPosition: Point = {
      x: tcPoint.x,
      y: tcPoint.y,
    };

    tcPoint.x += 250;

    nodes.push({
      id: tcId,
      data: { label: tcLabel },
      position: tcPosition,
    });

    const tcEdges = tc?.edges;

    for (const prop in tcEdges) {
      if (prop === "relations") {
        const relations = tcEdges[prop];
        let relPoint: Point = {
          x: tcPoint.x - 250,
          y: tcPoint.y + 250,
        };

        relations.forEach((rel, i) => {
          const relId = `${tcId}/rel/${rel.name}`;
          const relLabel = rel.name;
          const relPosition: Point = {
            x: relPoint.x,
            y: relPoint.y,
          };

          relPoint.x += 250;

          edges.push({
            id: `${tcId}-${relId}`,
            source: tcId,
            label: "relation",
            labelBgPadding: [8, 4],
            labelBgBorderRadius: 4,
            labelBgStyle: { fill: "#FFCC00", color: "#fff", fillOpacity: 0.7 },
            markerEnd: {
              type: MarkerType.ArrowClosed,
            },
            target: relId,
          });

          nodes.push({
            id: relId,
            data: { label: relLabel },
            position: relPosition,
          });
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
  const parsedData = getNodes(data);

  const [nodes, setNodes, onNodesChange] = useNodesState(parsedData.nodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(parsedData.edges);

  useEffect(() => {
    getNodes(data);
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
