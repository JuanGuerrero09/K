import { Icon, Popup } from "leaflet";
import { Marker } from "react-leaflet";

const Markers = ({ places }) => {
  const customIcon = new Icon({
    iconUrl: "https://www.svgrepo.com/show/532542/location-pin-lock.svg",
    iconSize: [20, 20],
  });
  const visitedIcon = new Icon({
    iconUrl: "./../images/visited.svg",
    iconSize: [20, 20],
  });
  const unvisitedIcon = new Icon({
    iconUrl: "https://www.svgrepo.com/show/532540/location-pin-alt-1.svg",
    iconSize: [20, 20],
  });

  function selectIcon(location) {
    const { isCompleted, isUnlocked } = location;
    if (isCompleted && isUnlocked) {
      return visitedIcon;
    }
    if (!isCompleted && isUnlocked) {
      return unvisitedIcon;
    }
    return customIcon;
  }

  console.log(places);

  return (
    <>
      {places.map((location, index) => {
        const { lat, lng } = location.geoCode;
        console.log(lat, lng);
        return (
          <Marker
            key={index}
            position={{ lat, lng }}
            icon={selectIcon(location)}
          >
            <Popup>
              <h2>{location.site_name}</h2>
            </Popup>
          </Marker>
        );
      })}
    </>
  );
};

export default Markers;
