import { Handle, Position } from "react-flow-renderer";

const PermissionNode = ({ data }: any) => {
  return (
    <div
      style={{
        backgroundColor: "#1acc92",
        padding: "1em",
        borderRadius: "10%",
        color: "#fefefe",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
        border: "3px solid #0c4633",
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

export default PermissionNode;
