import { Box, List, Tabs } from "@mantine/core";
import { LockAccess, ThreeDCubeSphere } from "tabler-icons-react";
import useTuples from "../hooks/useTuples";
import Tuple from "./Tuple";
import TupleAdd from "./TupleAdd";

interface Tuple {
  subject: {
    type: string;
    name: string;
  };
  relation: string;
  resource: {
    type: string;
    name: string;
  };
}

const TuplesArea = () => {
  const tuples = useTuples((s) => s.tuples);
  const subjects = useTuples((s) => s.getUniqueSubjects)();

  return (
    <Box
      style={{
        flex: 1,
        borderTop: "1px solid gray",
        overflow: "hidden",
        // height: "50%",
      }}
    >
      <Tabs
        style={{ marginTop: "0.5rem", height: "100%" }}
        styles={{ body: { height: "100%" } }}
        variant="default"
      >
        <Tabs.Tab label="Tuples" icon={<ThreeDCubeSphere size={14} />}>
          <Box style={{ overflow: "auto", height: "90%" }}>
            <List listStyleType={"none"}>
              <List.Item>
                <TupleAdd />
              </List.Item>
              {tuples.map((t, i) => (
                <List.Item key={i}>
                  <Tuple tuple={t} />
                </List.Item>
              ))}
            </List>
          </Box>
        </Tabs.Tab>
        <Tabs.Tab label="PermissionCheck" icon={<LockAccess size={14} />}>
          PermissionCheck
        </Tabs.Tab>
      </Tabs>
    </Box>
  );
};

export default TuplesArea;
