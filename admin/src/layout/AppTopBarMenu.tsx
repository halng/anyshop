"use client";
import React from "react";
import { Menubar } from "primereact/menubar";
import { InputText } from "primereact/inputtext";
import Link from "next/link";

const AppTopBarMenu = () => {
  const nestedMenuitems = [
    {
      label: "Customers",
      icon: "pi pi-fw pi-table",
      items: [
        {
          label: "New",
          icon: "pi pi-fw pi-user-plus",
          items: [
            {
              label: "Customer",
              icon: "pi pi-fw pi-plus",
            },
            {
              label: "Duplicate",
              icon: "pi pi-fw pi-copy",
            },
          ],
        },
        {
          label: "Edit",
          icon: "pi pi-fw pi-user-edit",
        },
      ],
    },
    {
      label: "Orders",
      icon: "pi pi-fw pi-shopping-cart",
      items: [
        {
          label: "View",
          icon: "pi pi-fw pi-list",
        },
        {
          label: "Search",
          icon: "pi pi-fw pi-search",
        },
      ],
    },
    {
      label: "Shipments",
      icon: "pi pi-fw pi-envelope",
      items: [
        {
          label: "Tracker",
          icon: "pi pi-fw pi-compass",
        },
        {
          label: "Map",
          icon: "pi pi-fw pi-map-marker",
        },
        {
          label: "Manage",
          icon: "pi pi-fw pi-pencil",
        },
      ],
    },
    {
      label: "Profile",
      icon: "pi pi-fw pi-user",
      items: [
        {
          label: "Settings",
          icon: "pi pi-fw pi-cog",
        },
        {
          label: "Billing",
          icon: "pi pi-fw pi-file",
        },
      ],
    },
    {
      label: "Quit",
      icon: "pi pi-fw pi-sign-out",
    },
  ];

  const menubarEndTemplate = () => {
    return (
      <span className="p-input-icon-left">
        <i className="pi pi-search" />
        <InputText type="text" placeholder="Search" />
      </span>
    );
  };

  const menubarStartTemplate = () => {
    return (
      <Link href="/" className="layout-topbar-logo">
        <img
          src={`/layout/images/anyshop.svg`}
          width="47.22px"
          height={"35px"}
          alt="logo"
        />
        <span>ANYSHOP</span>
      </Link>
    );
  };
  return (
    <Menubar
      className="d-flex w-100"
      model={nestedMenuitems}
      start={menubarStartTemplate}
      end={menubarEndTemplate}
    ></Menubar>
  );
};

export default AppTopBarMenu;
