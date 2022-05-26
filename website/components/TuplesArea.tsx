import { Box, Container, List, Tabs } from "@mantine/core";
import { useState } from "react";
import { LockAccess, ThreeDCubeSphere } from "tabler-icons-react";
import Tuple from "./Tuple";

const initialTuples = [
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
  const [tuples, setTuples] = useState(initialTuples);

  console.log(tuples);

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
          <Container style={{ overflow: "auto", height: "90%" }}>
            <List listStyleType={"none"}>
              {tuples.map((t) => (
                <List.Item>
                  <Tuple
                    subject={{ type: t.subject.type, name: t.subject.name }}
                    relation={t.relation}
                    resource={{ type: t.resource.type, name: t.resource.name }}
                    setTuples={setTuples}
                  />
                </List.Item>
              ))}
            </List>
          </Container>
        </Tabs.Tab>
        <Tabs.Tab label="PermissionCheck" icon={<LockAccess size={14} />}>
          PermissionCheck
        </Tabs.Tab>
      </Tabs>
    </Box>
  );
};

export default TuplesArea;
