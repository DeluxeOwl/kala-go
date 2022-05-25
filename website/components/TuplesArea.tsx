import { Box, Tabs } from "@mantine/core";
import { LockAccess, ThreeDCubeSphere } from "tabler-icons-react";

const TuplesArea = () => {
  return (
    <Box
      style={{
        flex: 1,
        borderTop: "1px solid gray",
      }}
    >
      <Tabs style={{ marginTop: "0.5rem" }} variant="default">
        <Tabs.Tab label="Tuples" icon={<ThreeDCubeSphere size={14} />}>
          Tuples
        </Tabs.Tab>
        <Tabs.Tab label="PermissionCheck" icon={<LockAccess size={14} />}>
          PermissionCheck
        </Tabs.Tab>
      </Tabs>
    </Box>
  );
};

export default TuplesArea;
