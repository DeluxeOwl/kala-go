import { Title } from "@mantine/core";
import { Handle, Position } from "react-flow-renderer";

const SubjectNode = ({ data }: any) => {
  return (
    <div
      style={{
        backgroundColor: "#b22cd3",
        padding: "1em",
        borderRadius: "10px",
        color: "#fefefe",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Handle
        type="target"
        position={Position.Top}
        id={`${data.id}.top`}
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

export default SubjectNode;
