import { Box } from "@mantine/core";

type GraphAreaProps = {
  children?: React.ReactNode;
};

const GraphArea = ({ children }: GraphAreaProps) => {
  return <Box style={{ flex: 4, height: "100%" }}>{children}</Box>;
};

export default GraphArea;
