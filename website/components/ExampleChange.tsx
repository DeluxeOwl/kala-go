import { Anchor, Group, Select, Text, ThemeIcon } from "@mantine/core";
import Image from "next/image";
import { forwardRef } from "react";
import {
  BarrierBlock,
  BrandGithub,
  BrandGoogleDrive,
  Folder,
  LockAccess,
  User,
} from "tabler-icons-react";
import logo from "../public/logo.png";

const data = [
  {
    icon: <User />,
    label: "Your Authorization model",
    value: "Your Authorization model",
    description: "Your own custom authorization model",
  },
  // {
  //   icon: <BrandGithub />,
  //   label: "Github",
  //   value: "Github",
  //   description: "Github's organization and member permission",
  // },
  {
    icon: <LockAccess />,
    label: "RBAC",
    value: "RBAC",
    description: "Example of simple RBAC",
  },
  {
    icon: <Folder />,
    label: "Documents (default)",
    value: "Documents",
    description: "Simple example showcasing all the posibilites",
  },
  {
    icon: <BrandGoogleDrive />,
    label: "Google Drive",
    value: "Google Drive",
    description: "GDrive's permissions for folders and documents",
  },
  {
    icon: <BarrierBlock />,
    label: "Blocklist",
    value: "Blocklist",
    description: "Example of a blocklist with banned user",
  },
  // {
  //   icon: <Users />,
  //   label: "Custom Roles",
  //   value: "Custom Roles",
  //   description: "Example of custom roles",
  // },
];

interface ItemProps extends React.ComponentPropsWithoutRef<"div"> {
  icon: React.ReactNode;
  label: string;
  description: string;
}
const SelectItem = forwardRef<HTMLDivElement, ItemProps>(
  ({ icon, label, description, ...others }: ItemProps, ref) => (
    <div ref={ref} {...others}>
      <Group noWrap>
        {icon}

        <div>
          <Text size="sm">{label}</Text>
          <Text size="xs" color="dimmed">
            {description}
          </Text>
        </div>
      </Group>
    </div>
  )
);

SelectItem.displayName = "SelectItem";

interface ExampleChangeProps {
  value: string;
  setValue: any;
}

export default function ExampleChange({ value, setValue }: ExampleChangeProps) {
  return (
    <Group
      style={{
        zIndex: 100,
        position: "relative",
        right: 18,
        backgroundColor: "transparent",
        boxShadow: "none",
      }}
      position="center"
    >
      <Image alt="logo-gopher" width={"45px"} height={"45px"} src={logo} />
      <Anchor href="https://github.com/DeluxeOwl/kala-go" target={"_blank"}>
        {" "}
        <ThemeIcon radius={"xl"} size="lg" color={"gray"}>
          <BrandGithub />
        </ThemeIcon>
      </Anchor>
      <Select
        placeholder="Pick an authorization model"
        itemComponent={SelectItem}
        data={data}
        maxDropdownHeight={400}
        radius="lg"
        value={value}
        onChange={setValue}
      />
    </Group>
  );
}
