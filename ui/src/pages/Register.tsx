import {
  Paper,
  createStyles,
  TextInput,
  PasswordInput,
  Checkbox,
  Button,
  Title,
  rem,
  Notification,
} from "@mantine/core";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useRegister } from "../hooks/useRegister";

const useStyles = createStyles((theme) => ({
  wrapper: {
    minHeight: rem(900),
    backgroundSize: "cover",
    backgroundImage:
      "url(https://images.unsplash.com/photo-1484242857719-4b9144542727?ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&ixlib=rb-1.2.1&auto=format&fit=crop&w=1280&q=80)",
  },

  form: {
    borderRight: `${rem(1)} solid ${
      theme.colorScheme === "dark" ? theme.colors.dark[7] : theme.colors.gray[3]
    }`,
    minHeight: rem(900),
    maxWidth: rem(450),
    paddingTop: rem(80),

    [theme.fn.smallerThan("sm")]: {
      maxWidth: "100%",
    },
  },

  title: {
    color: theme.colorScheme === "dark" ? theme.white : theme.black,
    fontFamily: `Greycliff CF, ${theme.fontFamily}`,
  },
}));

const RegisterPage = ({}) => {
  const { classes } = useStyles();
  const navigate = useNavigate();
  const [formState, setFormState] = useState({
    username: "",
    email: "",
    password: "",
  });
  const [notificationState, setNotificationState] = useState({
    color: "",
    title: "",
    content: "",
    visible: false,
  });
  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setFormState((old) => ({
      ...old,
      [event.target.name]: event.target.value,
    }));
  };
  const { isLoading, mutate } = useRegister({
    onError: () => {
      setNotificationState({
        color: "red",
        title: "register error",
        content: "",
        visible: true,
      });
    },
    onSuccess: () => {
      setNotificationState({
        color: "green",
        title: "registered successfully",
        content: "",
        visible: true,
      });
      navigate("/requests");
    },
  });
  return (
    <div className={classes.wrapper}>
      <Paper className={classes.form} radius={0} p={30}>
        <Title order={2} className={classes.title} ta="center" mt="md" mb={50}>
          Register to use the Citizen Request app!
        </Title>

        <TextInput
          name="username"
          onChange={handleChange}
          label="username"
          placeholder="jon Doe"
          size="md"
        />
        <TextInput
          label="Email address"
          name="email"
          onChange={handleChange}
          placeholder="hello@gmail.com"
          size="md"
        />
        <PasswordInput
          name="password"
          onChange={handleChange}
          label="Password"
          placeholder="Your password"
          mt="md"
          size="md"
        />
        <Checkbox label="Keep me logged in" mt="xl" size="md" />
        <Button
          loading={isLoading}
          onClick={() => {
            mutate(formState);
          }}
          fullWidth
          mt="xl"
          size="md"
        >
          Register
        </Button>
        {notificationState.visible && (
          <Notification
            title={notificationState.title}
            color={notificationState.color}
          >
            {notificationState.content}
          </Notification>
        )}
      </Paper>
    </div>
  );
};

export default RegisterPage;
