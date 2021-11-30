import React from "react";
import { render, screen } from "@testing-library/react";

import { Header } from "../Header";

describe("Header", () => {
  it("should have a heading", () => {
    render(<Header />);
    const heading = screen.getByRole("heading", { level: 1 });

    expect(heading).toHaveTextContent("todos");
  });
});
