import { Box } from "@mantine/core";
import useGraph from "../hooks/useGraph";
import ErrorGraph from "./ErrorGraph";
import Graph from "./Graph";
import LoadingGraph from "./LoadingGraph";

type GraphAreaProps = {
  children?: React.ReactNode;
};

const GraphArea = ({ children }: GraphAreaProps) => {
  const { isLoading, isError, data, error } = useGraph();
  console.log(data);

  // @ts-ignore
  const errorMessage: string = error?.message;

  if (isLoading) {
    return <LoadingGraph />;
  }

  if (isError) {
    return <ErrorGraph errorMessage={errorMessage} />;
  }

  return (
    <Box style={{ flex: 4, height: "100%", overflow: "auto" }}>
      <Graph data={data} />
    </Box>
  );
};

export default GraphArea;
