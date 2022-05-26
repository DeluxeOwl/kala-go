import {
  Box,
  Container,
  Group,
  Kbd,
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
};

export default function Tuple(props: TupleProps) {
  return (
    <Container
      style={{
        border: "1px solid gray",
        margin: "1rem",
        borderRadius: "5px",
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
          <Box>{`${props.subject.type}:${props.subject.name}`}</Box>
          <Box>{`${props.relation}`}</Box>
          <Box>{`${props.resource.type}:${props.resource.name}`}</Box>
        </Stack>

        <Stack style={{ flex: 1 }}>
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
              radius={"lg"}
              variant="gradient"
              gradient={{ from: "teal", to: "blue", deg: 60 }}
            >
              <QuestionMark />
            </ThemeIcon>
          </Tooltip>
        </Stack>
      </Group>
    </Container>
  );
}
