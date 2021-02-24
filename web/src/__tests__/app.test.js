import React from "react";
import axios from "axios";
import Enzyme, { shallow } from "enzyme";
import {
  render,
  fireEvent,
  getByTestId,
  act,
  waitFor,
} from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";
import App from "../app.js";

afterEach(() => {
  window.localStorage.removeItem("user_token");
});

jest.mock("axios");

it("renders dashboard after login", () => {
  const component = shallow(<App />);

  const button = component.find("button");
  expect(button.text()).toEqual("Login to my Dashboard");
});

it("starts logged out", async () => {
  const { container } = render(<App />);
  const loginBtn = getByTestId(container, "login-button");
  const inputEmail = getByTestId(container, "email");
  const inputPassword = getByTestId(container, "password");

  const loginResult = {
    status: 200,
    data: {
      access_token:
        "79058f1d814c80a50a197d5faea4cada8f65b9fa2bbbaf793bc5dfcc89cb0961",
      account_id: "47f3c307-6344-49e7-961c-ea200e950a89",
    },
  };
  axios.post.mockImplementationOnce(() => Promise.resolve(loginResult));

  await act(async () => {
    fireEvent.change(inputEmail, { target: { value: "test@example.com" } });
    fireEvent.change(inputPassword, { target: { value: "testpassword" } });
    fireEvent.click(loginBtn);
  });

  expect(window.localStorage.getItem("user_token")).toBe(
    loginResult.data.access_token
  );
  await waitFor(() => {
    const title = getByTestId(container, "dashboard-title");
    expect(title).toHaveTextContent("User Management Dashboard");
  });
});
