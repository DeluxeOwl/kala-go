import { Kbd } from "@mantine/core";
import { showNotification } from "@mantine/notifications";
import { Help, X } from "tabler-icons-react";

const showError = (errorMessage: string) => {
  showNotification({
    title: "Error",
    icon: <X size={18} />,
    color: "red",
    message: errorMessage,
    radius: "lg",
    style: { whiteSpace: "pre-line" },
  });
};

const showHelpNotif = () => {
  showNotification({
    title: "Note",
    icon: <Help size={18} />,
    color: "green",
    message: (
      <div>
        You can press
        <Kbd>Ctrl</Kbd> + <Kbd>Shift</Kbd> + <Kbd>K</Kbd> at any time to bring
        up the help menu.
      </div>
    ),
    radius: "lg",
    style: { whiteSpace: "pre-line" },
  });
};

export { showError, showHelpNotif };
