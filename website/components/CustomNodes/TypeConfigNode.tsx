import { Title } from "@mantine/core";
import { Handle, Position } from "react-flow-renderer";

const TypeConfigNode = ({ data }: any) => {
  return (
    <div
      style={{
        backgroundColor: "#8787ff",
        padding: "1em",
        borderRadius: "50%",
        color: "#fefefe",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
        border: "3px solid #6262fc",
      }}
    >
      <Handle
        type="target"
        position={Position.Bottom}
        id={`${data.id}.bottom`}
        style={{ borderRadius: "0", visibility: "hidden" }}
      />
      <div id={data.id} style={{ margin: "1em" }}>
        <Title order={3}>{data.label}</Title>
      </div>
      <Handle
        type="source"
        position={Position.Right}
        id={`${data.id}.right1`}
        style={{ top: "30%", borderRadius: 0, visibility: "hidden" }}
      />
    </div>
  );
};

export default TypeConfigNode;
