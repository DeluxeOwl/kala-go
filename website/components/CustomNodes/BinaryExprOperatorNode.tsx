import { Handle, Position } from "react-flow-renderer";

const BinaryExprOperatorNode = ({ data }: any) => {
  let colorDependingOnLabel: string;
  switch (data.label) {
    case "|":
      colorDependingOnLabel = "#008000";
      break;

    case "&":
      colorDependingOnLabel = "#f29d02";
      break;

    case "!":
      colorDependingOnLabel = "#ff0000";
      break;

    default:
      colorDependingOnLabel = "#fff";
      break;
  }

  return (
    <div
      style={{
        backgroundColor: colorDependingOnLabel,
        padding: "1em",
        width: "50px",
        height: "50px",
        borderRadius: "5%",
        color: "#fefefe",
        display: "inline-flex",
        alignItems: "center",
        justifyContent: "center",
        border: `3px solid ${colorDependingOnLabel}`,
      }}
    >
      <Handle
        type="target"
        position={Position.Top}
        id={`${data.id}.top`}
        style={{ borderRadius: "0", visibility: "hidden" }}
      />

      <div id={data.id}>{data.label}</div>
      <Handle
        type="source"
        position={Position.Bottom}
        id={`${data.id}.bottom`}
        style={{ top: "30%", borderRadius: 0, visibility: "hidden" }}
      />
    </div>
  );
};

export default BinaryExprOperatorNode;
