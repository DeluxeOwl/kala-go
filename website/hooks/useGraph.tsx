import { useQuery } from "react-query";
import { BACKEND_URL } from "../url";

const fetchGraph = () =>
  fetch(`${BACKEND_URL}/graph`).then((res) => res.json());

function useGraph() {
  const { isLoading, isError, data, error, refetch } = useQuery(
    "graph",
    fetchGraph
  );

  return { isLoading, isError, data, error, refetch };
}

export default useGraph;
