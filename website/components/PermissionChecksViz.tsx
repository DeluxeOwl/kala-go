import { Button, Group, Text, ThemeIcon } from "@mantine/core";
import { CircleCheck, CircleOff, PlayerTrackNext } from "tabler-icons-react";
export default function PermissionChecksViz() {
  return (
    <Group>
      <Button
        rightIcon={<PlayerTrackNext />}
        variant="subtle"
        size="lg"
        radius="xl"
        color="violet"
      >
        Run all checks
      </Button>
      <ThemeIcon color="green" size="lg" radius="lg">
        <CircleCheck />
      </ThemeIcon>
      <Text color="green" size="lg">
        0
      </Text>
      <ThemeIcon color="red" size="lg" radius="lg">
        <CircleOff />
      </ThemeIcon>
      <Text color="red" size="lg">
        0
      </Text>
    </Group>
  );
}
