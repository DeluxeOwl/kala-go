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

type TupleProps = {
  subject: { type: string; name: string };
  relation: string;
  resource: { type: string; name: string };
  setTuples: CallableFunction;
};

export default function Tuple(props: TupleProps) {
  const handleTupleDelete = () => {
    props.setTuples((tuples: any) =>
      tuples.filter(
        (t: any) =>
          !(
            t.subject.name === props.subject.name &&
            t.subject.type === props.subject.type &&
            t.relation === props.relation &&
            t.resource.name === props.resource.name &&
            t.resource.type === props.resource.type
          )
      )
    );
  };

  return (
    <Container
      style={{
        border: "2px solid gray",
        margin: "1rem",
        borderRadius: "10px",
        padding: "5px",
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
            <Mark color="indigo">{`${props.subject.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${props.subject.name}`}</Mark>
          </Box>
          <Box>
            <Badge color="cyan" size="md">{`${props.relation}`}</Badge>
          </Box>
          <Box>
            <Mark color="indigo">{`${props.resource.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${props.resource.name}`}</Mark>
          </Box>
        </Stack>

        <Group style={{ flex: 1 }}>
          <Tooltip
            label={
              <>
                <Kbd>{`${props.subject.type}:${props.subject.name}`}</Kbd>
                &nbsp;is a&nbsp;<Kbd>{`${props.relation}`}</Kbd>
                &nbsp;of&nbsp;
                <Kbd>{`${props.resource.type}:${props.resource.name}`}</Kbd>
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
