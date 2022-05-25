import { Box, Tabs } from "@mantine/core";
import { LockAccess, ThreeDCubeSphere } from "tabler-icons-react";

const tuples = [
  {
    subject: {
      type: "user",
      name: "anna",
    },
    relation: "reader",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "anna",
    },
    relation: "writer",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "folder",
      name: "secret_folder",
    },
    relation: "parent_folder",
    resource: {
      type: "document",
      name: "report.csv",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "user",
      name: "john",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "dev",
    },
  },
  {
    subject: {
      type: "group",
      name: "dev#member",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "group",
      name: "test_group#member",
    },
    relation: "reader",
    resource: {
      type: "folder",
      name: "secret_folder",
    },
  },
  {
    subject: {
      type: "user",
      name: "steve",
    },
    relation: "member",
    resource: {
      type: "group",
      name: "dev",
    },
  },
];

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
          {/* <ScrollArea>
            <List listStyleType={"none"}>
              {tuples.map((t) => (
                <List.Item>
                  <Tuple
                    subject={{ type: t.subject.type, name: t.subject.name }}
                    relation={t.relation}
                    resource={{ type: t.resource.type, name: t.resource.name }}
                  />
                </List.Item>
              ))}
            </List>
          </ScrollArea> */}
        </Tabs.Tab>
        <Tabs.Tab label="PermissionCheck" icon={<LockAccess size={14} />}>
          PermissionCheck
        </Tabs.Tab>
      </Tabs>
    </Box>
  );
};

export default TuplesArea;
