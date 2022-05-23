import { Box } from "@mantine/core";
import useGraph from "../hooks/useGraph";
import Graph from "./Graph";
type GraphAreaProps = {
  children?: React.ReactNode;
};

const GraphArea = ({ children }: GraphAreaProps) => {
  const { isLoading, isError, data, error, refetch } = useGraph();

  if (isLoading) {
    return <Box style={{ flex: 4, height: "100%" }}>{"Loading ..."}</Box>;
  }

  if (isError) {
    // @ts-ignore
    return <Box style={{ flex: 4, height: "100%" }}>{error?.message}</Box>;
  }

  return (
    <Box style={{ flex: 4, height: "100%", overflow: "auto" }}>
      <Graph data={data} />
    </Box>
  );
};

export default GraphArea;