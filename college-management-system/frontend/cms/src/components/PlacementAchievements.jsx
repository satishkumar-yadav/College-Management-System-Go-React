// src/components/PlacementAchievements.jsx
import '../styles/placementAchievements.css';

const achievements = [
  { title: "Top 10 Engineering College (State Rank)", year: "2023" },
  { title: "Best Campus Award - AICTE", year: "2022" },
];

const placements = [
  { company: "TCS", students: 120, package: "6.5 LPA" },
  { company: "Infosys", students: 90, package: "5.8 LPA" },
];

const PlacementAchievements = () => (
  <section className="placement-section">
    <h2>🎖 Achievements</h2>
    <ul>
      {achievements.map((item, i) => (
        <li key={i}>{item.year} - {item.title}</li>
      ))}
    </ul>

    <h2>💼 Placements</h2>
    <table>
      <thead>
        <tr><th>Company</th><th>Students Placed</th><th>Package</th></tr>
      </thead>
      <tbody>
        {placements.map((item, i) => (
          <tr key={i}>
            <td>{item.company}</td>
            <td>{item.students}</td>
            <td>{item.package}</td>
          </tr>
        ))}
      </tbody>
    </table>
  </section>
);

export default PlacementAchievements;
