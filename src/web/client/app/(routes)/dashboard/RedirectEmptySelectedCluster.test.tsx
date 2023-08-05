import { render, act } from "@testing-library/react";
import RedirectEmptySelectedCluster from "./RedirectEmptySelectedCluster";
import "@testing-library/jest-dom/extend-expect";

describe("RedirectEmptySelectedCluster", () => {
  let setSelectedClusterIdMock = jest.fn();
  let SelectedClusterId = 1;
  // Mock the useAppState hook
  jest.mock("../../../store/appstate", () => ({
    useAppState: jest.fn().mockReturnValue({
      SelectedClusterId: SelectedClusterId,
      SetSelectedClusterId: setSelectedClusterIdMock,
    }),
  }));

  it("should not call SetSelectedClusterId when SelectedClusterId is undefined", () => {
    // Mock the useAppState hook return values
    const setSelectedClusterIdMock = jest.fn();
    jest.mock("../../../store/appstate", () => ({
      SelectedClusterId: undefined,
      SetSelectedClusterId: setSelectedClusterIdMock,
    }));

    // Render the component
    render(<RedirectEmptySelectedCluster />);

    // The useEffect is called after the component is mounted
    // so we need to wait for useEffect to be executed
    act(() => {
      jest.runAllTimers();
    });

    // Expect that SetSelectedClusterId was not called
    expect(setSelectedClusterIdMock).not.toHaveBeenCalled();
  });
});
