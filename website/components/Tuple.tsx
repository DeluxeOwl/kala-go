import {
  Container,
  Group,
  Kbd,
  Stack,
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
        <Container>
          <Stack>
            <Container>SUBJECT</Container>
            <Container>RELATION</Container>
            <Container>RESOURCE</Container>
          </Stack>
        </Container>
        <Container>
          <Stack align={"flex-start"}>
            <Container>
              {`${props.subject.type}:${props.subject.name}`}
            </Container>
            <Container>{`${props.relation}`}</Container>
            <Container>
              {`${props.resource.type}:${props.resource.name}`}
            </Container>
          </Stack>
        </Container>
        <Container>
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
        </Container>
      </Group>
    </Container>
  );
}
