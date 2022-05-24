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

const refValueDelim = " | ";
const refSubrelationDelim = "#";
const parentRelDelim = ".";

// TODO: calculate some stuff here to look good
const getNodes = (graph: any): Node[] => {
  let nodes: Node[] = [];
  let edges: Edge[] = [];

  // https://stackoverflow.com/questions/5300938/calculating-the-position-of-points-in-a-circle
  let radius = 500;
  let degrees: number = 6.28319;
  if (graph.length > 0) {
    degrees = 6.28319 / graph.length;
  }

  graph.forEach((tc, i) => {
    const tcId = `tc/${tc.name}`;
    const tcLabel = tc.name;
    const tcPosition: Point = {
      x: radius * Math.cos(degrees * (i + 1)) + 50,
      y: radius * Math.sin(degrees * (i + 1)) + 50,
    };

    // tcPoint.x += 250;

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
          x: tcPosition.x,
          y: tcPosition.y + 250,
        };

        relations.forEach((rel, i) => {
          const relId = `${tcId}/rel/${rel.name}`;
          const edgeId = `${tcId}-${relId}`;
          const relLabel = rel.name;
          const relPosition: Point = {
            x: relPoint.x,
            y: relPoint.y,
          };

          relPoint.x += 250;

          edges.push({
            id: edgeId,
            source: tcId,
            label: "relation",
            labelBgPadding: [8, 4],
            labelBgBorderRadius: 4,
            labelBgStyle: { fill: "#FFCC00", color: "#fff", fillOpacity: 0.7 },
            markerEnd: {
              type: MarkerType.ArrowClosed,
            },
            target: relId,
            style: {
              stroke: "yellow",
            },
          });

          nodes.push({
            id: relId,
            data: { label: relLabel },
            position: relPosition,
          });

          // Composed relation
          const relValue: string = rel?.value;
          if (relValue.includes(refValueDelim)) {
            // console.log(true, relValue);
          } else {
            edges.push({
              id: `${tcId}-${relId}-${relValue}`,
              source: relId,
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
              target: `tc/${relValue}`,
            });
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
