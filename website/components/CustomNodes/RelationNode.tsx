import { Handle, Position } from "react-flow-renderer";

const RelationNode = ({ data }: any) => {
  return (
    <div
      style={{
        backgroundColor: "#ffa887",
        padding: "1em",
        borderRadius: "25%",
        color: "#141414",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
        border: "3px solid #a85b3e",
      }}
    >
      <Handle
        type="target"
        position={Position.Left}
        id={`${data.id}.left`}
        style={{ borderRadius: "0", visibility: "hidden" }}
      />

      <div id={data.id}>{data.label}</div>
      <Handle
        type="source"
        position={Position.Right}
        id={`${data.id}.right1`}
        style={{ top: "30%", borderRadius: 0, visibility: "hidden" }}
      />
    </div>
  );
};

export default RelationNode;
