/* eslint-disable @next/next/no-img-element */
"use client";
import { useRouter } from "next/navigation";
import React, { useContext, useState } from "react";
import { Checkbox } from "primereact/checkbox";
import { Button } from "primereact/button";
import { Password } from "primereact/password";
import { LayoutContext } from "../../../layout/context/layoutcontext";
import { InputText } from "primereact/inputtext";
import { classNames } from "primereact/utils";
import { ShinyText } from "@components";
import { Register } from "@/types/auth";
import { AuthService } from "@/service";
import { Toast } from "primereact/toast";
const RegisterPage = () => {
  const [registerData, setRegisterData] = useState<Register>({
    email: "",
    password: "",
    confirmPassword: "",
    username: "",
  });
  const toast = React.useRef<Toast>(null);

  const [checked, setChecked] = useState(false);
  const { layoutConfig } = useContext(LayoutContext);

  const router = useRouter();
  const containerClassName = classNames(
    "surface-ground flex align-items-center justify-content-center min-h-screen min-w-screen overflow-hidden",
    { "p-input-filled": layoutConfig.inputStyle === "filled" }
  );

  const handleRegister = () => {
    // validate the form data
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(registerData.email)) {
      toast.current?.show({
        severity: "error",
        summary: "Invalid Email",
        detail: "Please enter a valid email address.",
      });
      setRegisterData({ ...registerData, email: "" });
      return;
    }

    if (registerData.password !== registerData.confirmPassword) {
      toast.current?.show({
        severity: "error",
        summary: "Password Mismatch",
        detail: "Passwords do not match.",
      });
      setRegisterData({
        ...registerData,
        password: "",
        confirmPassword: "",
      });
      return;
    }

    const usernameRegex = /^[a-zA-Z0-9_]{8,}$/;
    if (!usernameRegex.test(registerData.username)) {
      toast.current?.show({
        severity: "error",
        summary: "Invalid Username",
        detail:
          "Username must be at least 3 characters long and can only contain letters, numbers, and underscores.",
      });
      setRegisterData({ ...registerData, username: "" });
      return;
    }

    AuthService.register(registerData)
      .then(() => {
        toast.current?.show({
          severity: "success",
          summary: "Registration Successful",
          detail: "You can now log in with your credentials.",
        });
        // Simulate a successful registration
        setTimeout(() => {
          router.push("/login");
        }, 1000);
      })
      .catch((error) => {
        console.error("Registration failed:", error);
        toast.current?.show({
          severity: "error",
          summary: "Registration Failed",
          detail: "An error occurred during registration. Please try again.",
        });
      });
  };

  const isEnableRegisterButton =
    registerData.email !== "" &&
    registerData.password !== "" &&
    registerData.confirmPassword !== "" &&
    registerData.username !== "" &&
    registerData.password === registerData.confirmPassword &&
    checked;

  return (
    <div className={containerClassName}>
      <Toast ref={toast} />
      <div className="flex flex-column align-items-center justify-content-center">
        <img
          src={`/layout/images/anyshop.svg`}
          alt="ANYSHOP logo"
          className="mb-5 w-6rem flex-shrink-0"
        />
        <div
          style={{
            borderRadius: "56px",
            padding: "0.3rem",
            background:
              "linear-gradient(180deg, var(--primary-color) 10%, rgba(33, 150, 243, 0) 30%)",
          }}
        >
          <div
            className="w-full surface-card py-8 px-5 sm:px-8"
            style={{ borderRadius: "53px" }}
          >
            <div className="text-center mb-5">
              <ShinyText
                className="text-900 text-3xl font-medium mb-3 "
                disabled={false}
                speed={3}
                text="Hello, Welcome to ANYSHOP"
              />
              <br />
              <ShinyText
                className="text-600 font-medium"
                disabled={false}
                speed={3}
                text="Create your account to get started"
              />
            </div>

            <div>
              <label
                htmlFor="email"
                className="block text-900 text-xl font-medium mb-2"
              >
                Email
              </label>
              <InputText
                id="email"
                type="email"
                placeholder="Email address"
                className="w-full md:w-30rem mb-5"
                style={{ padding: "1rem" }}
                value={registerData.email}
                onChange={(e) =>
                  setRegisterData({ ...registerData, email: e.target.value })
                }
              />

              <label
                htmlFor="username"
                className="block text-900 text-xl font-medium mb-2"
              >
                Username
              </label>
              <InputText
                id="username"
                type="text"
                placeholder="Username"
                className="w-full md:w-30rem mb-5"
                style={{ padding: "1rem" }}
                value={registerData.username}
                onChange={(e) =>
                  setRegisterData({ ...registerData, username: e.target.value })
                }
              />

              <label
                htmlFor="password"
                className="block text-900 font-medium text-xl mb-2"
              >
                Password
              </label>
              <Password
                inputId="password"
                value={registerData.password}
                onChange={(e) =>
                  setRegisterData({ ...registerData, password: e.target.value })
                }
                placeholder="Password"
                toggleMask
                className="w-full mb-5"
                inputClassName="w-full p-3 md:w-30rem"
              ></Password>

              <label
                htmlFor="confirm-password"
                className="block text-900 font-medium text-xl mb-2"
              >
                Confirm Password
              </label>
              <Password
                inputId="confirm-password"
                value={registerData.confirmPassword}
                onChange={(e) =>
                  setRegisterData({
                    ...registerData,
                    confirmPassword: e.target.value,
                  })
                }
                placeholder="Password"
                toggleMask
                className="w-full mb-5"
                inputClassName="w-full p-3 md:w-30rem"
              ></Password>

              <div className="flex align-items-center justify-content-between mb-5 gap-5">
                <div className="flex align-items-center">
                  <Checkbox
                    inputId="terms"
                    checked={checked}
                    onChange={(e) => setChecked(e.checked ?? false)}
                    className="mr-2"
                  ></Checkbox>
                  <label htmlFor="terms">
                    I agree to the terms and conditions
                  </label>
                </div>
              </div>

              <Button
                label="Register"
                className="w-full p-3 text-xl"
                onClick={handleRegister}
                disabled={!isEnableRegisterButton}
              ></Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default RegisterPage;
