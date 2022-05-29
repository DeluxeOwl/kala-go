import { Box, List, Tabs } from "@mantine/core";
import { LockAccess, ThreeDCubeSphere } from "tabler-icons-react";
import useTuples from "../hooks/useTuples";
import PermissionCheck from "./PermissionCheck";
import PermissionCheckAdd from "./PermissionCheckAdd";
import PermissionChecksViz from "./PermissionChecksViz";
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
  const pcs = useTuples((s) => s.permissionChecks);

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
        style={{ marginTop: "1rem", height: "100%" }}
        styles={(theme) => ({
          body: { height: "100%" },
          tabControl: { fontSize: theme.fontSizes.xl },
        })}
        variant="default"
        color={"violet"}
        tabPadding="lg"
      >
        <Tabs.Tab label="Tuples" icon={<ThreeDCubeSphere size={14} />}>
          <Box style={{ overflow: "auto", height: "85%" }}>
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
          <Box style={{ overflow: "auto", height: "85%" }}>
            <PermissionChecksViz />
            <List listStyleType={"none"}>
              <List.Item>
                <PermissionCheckAdd />
              </List.Item>
              {pcs.map((pc, i) => (
                <List.Item key={i}>
                  <PermissionCheck pc={pc} />
                </List.Item>
              ))}
            </List>
          </Box>
        </Tabs.Tab>
      </Tabs>
    </Box>
  );
};

export default TuplesArea;
