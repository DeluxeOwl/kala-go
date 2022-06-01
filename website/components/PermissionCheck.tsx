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
  UnstyledButton,
} from "@mantine/core";
import {
  CircleCheck,
  CircleOff,
  PlayerPlay,
  QuestionMark,
} from "tabler-icons-react";
import useTuples from "../hooks/useTuples";
import { default as IPc } from "../types/permissionCheck";
import { BACKEND_URL } from "../url";
import { postReq } from "../util/fetchAll";
import { showError } from "../util/notifications";

type PermissionCheckProps = {
  pc: IPc;
};

function sleep(time: number) {
  return new Promise((resolve) => setTimeout(resolve, time));
}

export default function PermissionCheck({ pc: pc }: PermissionCheckProps) {
  const removePermissionCheck = useTuples((s) => s.removePermissionCheck);
  const updatePcStatus = useTuples((s) => s.updatePermissionStatus);

  const handlePcDelete = () => {
    removePermissionCheck(pc);
  };

  const handlePermCheck = async () => {
    updatePcStatus(pc, undefined);

    await sleep(300);

    try {
      const res = await postReq(`${BACKEND_URL}/permission-check`, pc);
      const resJson = await res.json();
      updatePcStatus(pc, resJson?.permission);
    } catch (error) {
      // @ts-ignore
      showError(error.message);
    }
  };

  return (
    <Container
      style={{
        border: "2px solid gray",
        marginTop: "1rem",
        borderRadius: "10px",
        padding: "10px",
      }}
    >
      <Group>
        <Stack style={{ flex: 1 }} align="center">
          {pc.hasPermission === undefined && (
            <Tooltip label="Run this check to see the permission">
              <ThemeIcon radius="lg" color="gray">
                <QuestionMark size="lg" />
              </ThemeIcon>
            </Tooltip>
          )}
          {pc.hasPermission === true && (
            <Tooltip
              label={
                <>
                  <Kbd>{`${pc.subject.type}:${pc.subject.name}`}</Kbd>
                  &nbsp;has &nbsp;<Kbd>{`${pc.permission}`}</Kbd>
                  &nbsp;on &nbsp;
                  <Kbd>{`${pc.resource.type}:${pc.resource.name}`}</Kbd>
                </>
              }
            >
              <ThemeIcon radius="lg" color="green">
                <CircleCheck size="lg" />
              </ThemeIcon>
            </Tooltip>
          )}
          {pc.hasPermission === false && (
            <Tooltip
              label={
                <>
                  <Kbd>{`${pc.subject.type}:${pc.subject.name}`}</Kbd>
                  &nbsp;doesn&apos;t have &nbsp;<Kbd>{`${pc.permission}`}</Kbd>
                  &nbsp;on &nbsp;
                  <Kbd>{`${pc.resource.type}:${pc.resource.name}`}</Kbd>
                </>
              }
            >
              <ThemeIcon radius="lg" color="red">
                <CircleOff size="lg" />
              </ThemeIcon>
            </Tooltip>
          )}
        </Stack>
        <Stack style={{ flex: 1 }}>
          <Container>
            <Text weight={700}>SUBJECT</Text>
          </Container>
          <Container>
            <Text weight={700}>PERMISSION</Text>
          </Container>
          <Container>
            <Text weight={700}>RESOURCE</Text>
          </Container>
        </Stack>

        <Stack style={{ flex: 6 }}>
          <Box>
            <Mark color="indigo">{`${pc.subject.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${pc.subject.name}`}</Mark>
          </Box>
          <Box>
            <Badge
              color="violet"
              size="md"
              variant="filled"
            >{`${pc.permission}`}</Badge>
          </Box>
          <Box>
            <Mark color="indigo">{`${pc.resource.type}: `}</Mark>
            &nbsp;
            <Mark color="teal">{`${pc.resource.name}`}</Mark>
          </Box>
        </Stack>

        <Group style={{ flex: 1 }} align="center">
          <Tooltip
            label={
              <>
                Does <Kbd>{`${pc.subject.type}:${pc.subject.name}`}</Kbd>
                &nbsp;have &nbsp;<Kbd>{`${pc.permission}`}</Kbd>
                &nbsp;permission on &nbsp;
                <Kbd>{`${pc.resource.type}:${pc.resource.name}`}</Kbd> ?
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
            aria-label="Delete permission check"
            variant="light"
            color={"red"}
            radius="md"
            mb={5}
            onClick={handlePcDelete}
          />
          <UnstyledButton onClick={handlePermCheck}>
            <ThemeIcon
              radius={"md"}
              variant="gradient"
              gradient={{ from: "green", to: "cyan", deg: 60 }}
            >
              <PlayerPlay />
            </ThemeIcon>
          </UnstyledButton>
        </Group>
      </Group>
    </Container>
  );
}
