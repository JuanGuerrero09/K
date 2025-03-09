import "./../../styles.css";
import "leaflet/dist/leaflet.css";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import { Icon } from "leaflet";
import MarkerClusterGroup from "react-leaflet-cluster";
import useLocations from "../../hooks/useLocations";

export default function Map() {
  const locations = useLocations();
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
    let isCompleted = location.is_complete;
    console.log(location);
    let isUnlocked = location.is_unlocked;
    if (isCompleted && isUnlocked) {
      return visitedIcon;
    }
    if (!isCompleted && isUnlocked) {
      return unvisitedIcon;
    } else {
      return customIcon;
    }
  }

  return (
    <MapContainer center={[40.416775, -3.70379]} zoom={13.5}>
      {/* OPEN STREEN MAPS TILES */}
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      <MarkerClusterGroup>
        {locations.map((location) => {
          return (
            <Marker position={location.geoCode} icon={selectIcon(location)}>
              <Popup>
                <h2>{location.site_name}</h2>
              </Popup>
            </Marker>
          );
        })}
      </MarkerClusterGroup>
    </MapContainer>
  );
}
