import { Box } from "@mantine/core";

type GraphAreaProps = {
  children?: React.ReactNode;
};

const GraphArea = ({ children }: GraphAreaProps) => {
  return (
    <Box style={{ border: "1px solid red", flex: 2, height: "100%" }}>
      {children}
    </Box>
  );
};

export default GraphArea;
