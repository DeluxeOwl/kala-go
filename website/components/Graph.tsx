import { Checkbox, CheckboxGroup, Stack } from "@mantine/core";
import jsep from "jsep";
import { nanoid } from "nanoid";
import { useEffect, useMemo, useRef, useState } from "react";
import ReactFlow, {
  Background,
  Controls,
  Edge,
  Node,
  useEdgesState,
  useNodesState,
} from "react-flow-renderer";
import {
  binaryExprOperatorEdge,
  binaryExprOperatorNode,
  cosDegrees,
  memberExprEdge,
  permEdge,
  permNode,
  randomIntFromInterval,
  relComposedEdge,
  relComposedNode,
  relComposedSubrelEdge,
  relEdge,
  relNode,
  relToTcEdge,
  sinDegrees,
  subjectEdge,
  subjectNode,
  tcNode,
} from "../util/graphFunctions";
import BinaryExprOperatorNode from "./CustomNodes/BinaryExprOperatorNode";
import PermissionNode from "./CustomNodes/PermissionNode";
import RelationNode from "./CustomNodes/RelationNode";
import RelComposedNode from "./CustomNodes/RelComposedNode";
import TypeConfigNode from "./CustomNodes/TypeConfigNode";

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

// TODO: calculate some stuff here to look good
const getNodes = (graph: any): NodesAndEdges => {
  let nodes: Node[] = [];
  let edges: Edge[] = [];

  // https://stackoverflow.com/questions/5300938/calculating-the-position-of-points-in-a-circle
  let radius = 500;
  let degrees: number = 360;
  if (graph.length > 0) {
    degrees = 360 / graph.length;
  }

  graph.forEach((tc: any, i: number) => {
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

        relations.forEach((rel: any, i: number) => {
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

            // @ts-ignore
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

        const computedDgRel = degrees * (i + 1);

        let permPoint: Point = {
          x: tcPosition.x + 1.5 * radius * cosDegrees(computedDgRel),
          y: tcPosition.y + 1.5 * radius * sinDegrees(computedDgRel),
        };

        permissions.forEach((perm: any, i: number) => {
          const permId = `${tcId}/perm/${perm.name}`;
          const edgeId = `${tcId}-${permId}`;
          const permLabel = perm.name;
          const permPosition: Point = {
            x: permPoint.x,
            y: permPoint.y,
          };

          permPoint.y -= 350;
          permPoint.x -= 350;

          // edges.push(permEdge(edgeId, tcId, permId));

          nodes.push(permNode(permId, permLabel, permPosition));
          edges.push(permEdge(edgeId, tcId, permId));

          const treeRoot = jsep(perm.value);

          const recursivePermCreation = (
            node: jsep.Expression,
            sourceId: string
          ) => {
            switch (node.type) {
              case "Identifier":
                // Means direct relation
                if (sourceId.includes("operator/")) {
                  const operator = sourceId.split("operator/")[1];
                  edges.push(
                    binaryExprOperatorEdge(
                      `${permId}/computedTree/${nanoid(6)}`,
                      sourceId,
                      `${tcId}/rel/${node.name}`,
                      operator
                    )
                  );
                } else {
                  edges.push({
                    id: `${permId}/computedTree/${nanoid(6)}`,
                    source: sourceId,
                    target: `${tcId}/rel/${node.name}`,
                  });
                }

                break;
              case "UnaryExpression":
                const notNodeId = `${permId}/computedTree/${nanoid(6)}`;

                nodes.push(
                  binaryExprOperatorNode(notNodeId, "!", {
                    x: permPoint.x + 500,
                    y: permPoint.y + 250,
                  })
                );

                edges.push(
                  binaryExprOperatorEdge(
                    `${permId}/computedTree/${nanoid(6)}`,
                    sourceId,
                    notNodeId,
                    "!"
                  )
                );
                edges.push(
                  binaryExprOperatorEdge(
                    `${permId}/computedTree/${nanoid(6)}`,
                    notNodeId,
                    // @ts-ignore
                    `${tcId}/rel/${node?.argument?.name}`,

                    "!"
                  )
                );

                break;
              case "BinaryExpression":
                // Insert Operator node, recurse
                const operatorNodeId = `${permId}/computedTree/${nanoid(
                  6
                )}/operator/${node.operator}`;

                nodes.push(
                  // @ts-ignore
                  binaryExprOperatorNode(operatorNodeId, node.operator, {
                    x: permPoint.x + randomIntFromInterval(100, 350),
                    y: permPoint.y + randomIntFromInterval(100, 550),
                  })
                );

                edges.push(
                  binaryExprOperatorEdge(
                    `${permId}/computedTree/${nanoid(6)}`,
                    sourceId,
                    operatorNodeId,
                    // @ts-ignore
                    node.operator
                  )
                );

                // @ts-ignore
                recursivePermCreation(node.left, operatorNodeId);
                // @ts-ignore
                recursivePermCreation(node.right, operatorNodeId);
                break;
              case "MemberExpression":
                // To current relation

                edges.push(
                  memberExprEdge(
                    `${permId}/computedTree/${nanoid(6)}`,
                    sourceId,
                    // @ts-ignore
                    `${tcId}/rel/${node?.object?.name}`
                  )
                );

                const referencedType = perm.edges.relations.find(
                  // @ts-ignore
                  (rel) => rel.name === node?.object?.name
                ).value;

                // from current relation to the specified permission
                edges.push(
                  memberExprEdge(
                    `${permId}/computedTree/${nanoid(6)}`,
                    // @ts-ignore
                    `${tcId}/rel/${node?.object?.name}`,
                    // @ts-ignore
                    `tc/${referencedType}/rel/${node?.property?.name}`
                  )
                );

                break;
              default:
                break;
            }
          };

          recursivePermCreation(treeRoot, `${tcId}/perm/${perm.name}`);
        });
      }
      if (prop === "subjects") {
        const subjects = tcEdges[prop];
        const computedDgRel = degrees * (i + 1);

        let subjPoint: Point = {
          x: tcPosition.x + 2 * radius * cosDegrees(computedDgRel),
          y: tcPosition.y + 2 * radius * sinDegrees(computedDgRel),
        };

        subjects.forEach((subj: any, i: number) => {
          const subjId = `${tcId}/subj/${subj.name}`;
          const edgeId = `${tcId}-${subjId}`;

          const subjPosition: Point = {
            x: subjPoint.x,
            y: subjPoint.y,
          };

          subjPoint.y += 250;
          subjPoint.x += 250;

          nodes.push(subjectNode(subjId, subj.name, subjPosition));
          edges.push(subjectEdge(edgeId, tcId, subjId));
        });
      }
    }
  });

  const ne: NodesAndEdges = {
    nodes: nodes,
    edges: edges,
  };

  // @ts-ignore
  return ne;
};

type GraphProps = {
  data: any;
};

const Graph = ({ data }: GraphProps) => {
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  const initialNodes = useRef<Node[]>([]);
  const initialEdges = useRef<Edge[]>([]);

  const nodeTypes = useMemo(
    () => ({
      typeConfigNode: TypeConfigNode,
      relationNode: RelationNode,
      permissionNode: PermissionNode,
      relComposedNode: RelComposedNode,
      binaryExprOperatorNode: BinaryExprOperatorNode,
    }),
    []
  );

  const [checkboxValues, setCheckboxValues] = useState<string[]>([
    "includesRelEdges",
  ]);

  useEffect(() => {
    const parsedData = getNodes(data);
    setNodes(parsedData.nodes);
    setEdges(parsedData.edges);

    initialNodes.current = parsedData.nodes;
    initialEdges.current = parsedData.edges;

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  useEffect(() => {
    setNodes(initialNodes.current);
    setEdges(initialEdges.current);
  }, [checkboxValues, setNodes, setEdges]);

  useEffect(() => {
    if (!checkboxValues.includes("includesRelEdges")) {
      setNodes((no) =>
        no.filter((n) => !(n.data.label === "|" && n.id.includes("/or")))
      );
      setEdges((ed) =>
        ed.filter((e) => !(e.label === "includes" || e.label === "OR"))
      );
    }
    if (!checkboxValues.includes("includesSubjects")) {
      setNodes((no) => no.filter((n) => !n.id.includes("/subj/")));
      setEdges((ed) => ed.filter((e) => !e.id.includes("/subj/")));
    }
    if (!checkboxValues.includes("includesPermEdges")) {
      setNodes((no) => no.filter((n) => !n.id.includes("/computedTree/")));
      setEdges((ed) => ed.filter((e) => !e.id.includes("/computedTree/")));
    }
  }, [checkboxValues, setNodes, setEdges, data]);

  return (
    <ReactFlow
      nodes={nodes}
      onNodesChange={onNodesChange}
      edges={edges}
      onEdgesChange={onEdgesChange}
      nodesConnectable={false}
      // @ts-ignore
      connectionMode={"loose"}
      fitView
      snapToGrid={true}
      snapGrid={[20, 20]}
      nodeTypes={nodeTypes}
    >
      <Controls />
      <Background />
      <Stack style={{ position: "absolute", left: 10, right: 10, zIndex: 4 }}>
        <CheckboxGroup
          value={checkboxValues}
          onChange={setCheckboxValues}
          orientation="vertical"
          label="Select which nodes to show or hide"
          size="md"
        >
          <Checkbox
            value="includesRelEdges"
            label="Includes relations to type edges"
          />
          <Checkbox value="includesSubjects" label="Includes subject nodes" />
          <Checkbox
            value="includesPermEdges"
            label="Includes permission edges"
          />
        </CheckboxGroup>
      </Stack>
    </ReactFlow>
  );
};

export default Graph;
