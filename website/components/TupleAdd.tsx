import {
  Box,
  Button,
  Container,
  Group,
  Input,
  Stack,
  Text,
} from "@mantine/core";
import { useInputState } from "@mantine/hooks";
import { Plus } from "tabler-icons-react";
import useTuples from "../hooks/useTuples";

export default function TupleAdd() {
  const addTuple = useTuples((s) => s.addTuple);

  const [subjectType, setSubjectType] = useInputState("");
  const [subjectName, setSubjectName] = useInputState("");
  const [relation, setRelation] = useInputState("");
  const [resourceType, setResourceType] = useInputState("");
  const [resourceName, setResourceName] = useInputState("");

  const handleTupleAdd = () => {
    addTuple({
      subject: {
        type: subjectType,
        name: subjectName,
      },
      relation: relation,
      resource: {
        type: resourceType,
        name: resourceName,
      },
    });
  };

  return (
    <Container
      style={{
        border: "2px solid gray",
        margin: "1rem",
        borderRadius: "10px",
        padding: "10px",
      }}
    >
      <Group style={{ height: "100%" }}>
        <Stack style={{ flex: 1 }}>
          <Container>
            <Text weight={700}>SUBJECT</Text>
          </Container>
          <Container>
            <Text weight={700}>RELATION</Text>
          </Container>
          <Container>
            <Text weight={700}>RESOURCE</Text>
          </Container>
        </Stack>

        <Stack style={{ flex: 3 }}>
          <Input value={subjectType} onChange={setSubjectType} />
          <Input value={relation} onChange={setRelation} />
          <Input value={resourceType} onChange={setResourceType} />
        </Stack>

        <Group style={{ flex: 4, height: "100%" }}>
          <Stack style={{ flex: 2, height: "100%" }}>
            <Input
              value={subjectName}
              onChange={setSubjectName}
              style={{ margin: "auto" }}
            />
            <Box style={{ height: "36px" }} />
            <Input
              value={resourceName}
              onChange={setResourceName}
              style={{ margin: "auto" }}
            />
          </Stack>
          <Button
            style={{ flex: 1 }}
            variant="filled"
            color="green"
            rightIcon={<Plus size={16} />}
            onClick={handleTupleAdd}
          >
            Add tuple
          </Button>
        </Group>
      </Group>
    </Container>
  );
}
