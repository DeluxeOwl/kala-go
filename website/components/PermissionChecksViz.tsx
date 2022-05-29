import { Button, Group, Text, ThemeIcon } from "@mantine/core";
import {
  CircleCheck,
  CircleOff,
  PlayerTrackNext,
  QuestionMark,
} from "tabler-icons-react";
import useTuples from "../hooks/useTuples";
import { BACKEND_URL } from "../url";
import { postReq } from "../util/fetchAll";
export default function PermissionChecksViz() {
  const permissionChecks = useTuples((s) => s.permissionChecks);
  const updatePcStatus = useTuples((s) => s.updatePermissionStatus);

  const undefinedValues = permissionChecks.filter(
    (pc) => pc.hasPermission === undefined
  ).length;
  const falseValues = permissionChecks.filter(
    (pc) => pc.hasPermission === false
  ).length;
  const trueValues = permissionChecks.filter(
    (pc) => pc.hasPermission === true
  ).length;

  const handleAllRuns = async () => {
    permissionChecks.forEach(async (pc) => {
      try {
        const res = await postReq(`${BACKEND_URL}/permission-check`, pc);
        const resJson = await res.json();
        updatePcStatus(pc, resJson?.permission);
      } catch (error) {
        // @ts-ignore
        showError(error.message);
      }
    });
  };

  return (
    <Group
      style={{
        width: "50%",
        marginLeft: "auto",
        marginRight: "auto",
        marginBottom: "1rem",
      }}
      position="center"
    >
      <Button
        rightIcon={<PlayerTrackNext />}
        variant="subtle"
        size="lg"
        radius="xl"
        color="violet"
        onClick={handleAllRuns}
      >
        Run all checks
      </Button>
      <ThemeIcon color="green" size="lg" radius="lg" variant="outline">
        <CircleCheck />
      </ThemeIcon>
      <Text color="green" size="lg">
        {trueValues}
      </Text>
      <ThemeIcon color="red" size="lg" radius="lg" variant="outline">
        <CircleOff />
      </ThemeIcon>
      <Text color="red" size="lg">
        {falseValues}
      </Text>
      <ThemeIcon color="violet" size="lg" radius="lg" variant="outline">
        <QuestionMark />
      </ThemeIcon>
      <Text color="violet" size="lg">
        {undefinedValues}
      </Text>
    </Group>
  );
}
