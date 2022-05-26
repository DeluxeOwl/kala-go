import {
  Button,
  Container,
  Group,
  Input,
  Select,
  Stack,
  Text,
} from "@mantine/core";
import { Plus } from "tabler-icons-react";

type TupleAddProps = {
  setTuples: CallableFunction;
};

export default function TupleAdd(props: TupleAddProps) {
  const handleTupleAdd = () => {};

  return (
    <Container
      style={{
        border: "2px solid gray",
        margin: "1rem",
        borderRadius: "10px",
        padding: "5px",
      }}
    >
      <Group grow>
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

        <Stack style={{ flex: 6 }}>
          <Select data={[]} />
          <Select data={[]} />
          <Select data={[]} />
        </Stack>

        <Group style={{ flex: 1 }}>
          <Stack>
            <Input />
            <Input />
            <Input />
          </Stack>
          <Button variant="filled" color="green" rightIcon={<Plus size={16} />}>
            Add tuple
          </Button>
        </Group>
      </Group>
    </Container>
  );
}
