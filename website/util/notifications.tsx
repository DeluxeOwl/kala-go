import { Kbd, Text } from "@mantine/core";
import { showNotification } from "@mantine/notifications";
import { Help, X } from "tabler-icons-react";

const showError = (errorMessage: string) => {
  showNotification({
    title: <Text size="xl">Error</Text>,
    icon: <X size={18} />,
    color: "red",
    message: <Text size="md">{errorMessage}</Text>,
    radius: "lg",
    style: { whiteSpace: "pre-line" },
  });
};

const showHelpNotif = () => {
  showNotification({
    title: <Text size="xl">Note</Text>,
    icon: <Help size={18} />,
    color: "green",
    message: (
      <Text size="md">
        You can press
        <Kbd>Ctrl</Kbd> + <Kbd>Shift</Kbd> + <Kbd>K</Kbd> at any time to bring
        up the help menu.
      </Text>
    ),
    radius: "lg",
    style: { whiteSpace: "pre-line" },
  });
};

export { showError, showHelpNotif };
