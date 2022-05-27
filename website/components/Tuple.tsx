import {
  Badge,
  Box,
  CloseButton,
  Container,
  Group,
  Kbd,
  Mark,
  Stack,
  Text,
  ThemeIcon,
  Tooltip,
} from "@mantine/core";
import { QuestionMark } from "tabler-icons-react";
import useTuples from "../hooks/useTuples";
import { default as ITuple } from "../types/tuple";

type Tupletuple = {
  tuple: ITuple;
};

export default function Tuple({ tuple }: Tupletuple) {
  const removeTuple = useTuples((s) => s.removeTuple);

  const handleTupleDelete = () => {
    removeTuple(tuple);
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
      <Group>
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
          <Box>
            <Mark color="indigo">{`${tuple.subject.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${tuple.subject.name}`}</Mark>
          </Box>
          <Box>
            <Badge color="cyan" size="md">{`${tuple.relation}`}</Badge>
          </Box>
          <Box>
            <Mark color="indigo">{`${tuple.resource.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${tuple.resource.name}`}</Mark>
          </Box>
        </Stack>

        <Group style={{ flex: 1 }}>
          <Tooltip
            label={
              <>
                <Kbd>{`${tuple.subject.type}:${tuple.subject.name}`}</Kbd>
                &nbsp;is a&nbsp;<Kbd>{`${tuple.relation}`}</Kbd>
                &nbsp;of&nbsp;
                <Kbd>{`${tuple.resource.type}:${tuple.resource.name}`}</Kbd>
              </>
            }
            withArrow
            arrowSize={3}
          >
            <ThemeIcon
              radius={"md"}
              variant="gradient"
              gradient={{ from: "teal", to: "blue", deg: 60 }}
            >
              <QuestionMark />
            </ThemeIcon>
          </Tooltip>
          <CloseButton
            aria-label="Delete tuple"
            variant="light"
            color={"red"}
            radius="md"
            mb={5}
            onClick={handleTupleDelete}
          />
        </Group>
      </Group>
    </Container>
  );
}
