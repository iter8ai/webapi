import { renderToString } from "react-dom/server";
import { expect, test } from "vitest";
import { App } from "./App";

test("renders the page heading", () => {
  expect(renderToString(<App />)).toContain("Weather");
});
