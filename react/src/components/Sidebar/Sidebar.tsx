// import useLocations from "../../hooks/useLocations";

// const getCategories() => {
//
// }

export default function Sidebar() {
  return (
    <>
      <label
        htmlFor="my-drawer-4"
        aria-label="close sidebar"
        className="drawer-overlay"
      ></label>
      <ul className="menu flex bg-base-100 text-base-content min-h-full w-80 p-4">
        {/* Sidebar content here */}
        {/* For TSX uncomment the commented types below */}
        <h2 className="text-3xl mt-4 self-center font-bold text-blue-600">
          Progress
        </h2>
        <div
          className="radial-progress self-center mt-6"
          style={{ "--value": 70 } /* as React.CSSProperties */}
          aria-valuenow={70}
          role="progressbar"
        >
          70%
        </div>
        <li>
          <a>Sidebar Item 1</a>
        </li>
        <progress
          className="progress progress-primary w-56"
          value={50}
          max="100"
        ></progress>
        <li>
          <a>Sidebar Item 2</a>
        </li>
        <progress
          className="progress progress-primary w-56"
          value="70"
          max="100"
        ></progress>
      </ul>
    </>
  );
}
