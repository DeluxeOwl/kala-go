import { Title } from "@mantine/core";
import { Handle, Position } from "react-flow-renderer";

const RelComposedNode = ({ data }: any) => {
  return (
    <div
      style={{
        backgroundColor: "#008000",
        padding: "1em",
        width: "50px",
        height: "50px",
        borderRadius: "5%",
        color: "#fefefe",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
        border: "3px solid #008000",
      }}
    >
      <Handle
        type="target"
        position={Position.Top}
        id={`${data.id}.top`}
        style={{ borderRadius: "0", visibility: "hidden" }}
      />

      <div id={data.id}>
        <Title order={2}>{data.label}</Title>
      </div>
      <Handle
        type="source"
        position={Position.Bottom}
        id={`${data.id}.bottom`}
        style={{ top: "30%", borderRadius: 0, visibility: "hidden" }}
      />
    </div>
  );
};

export default RelComposedNode;
