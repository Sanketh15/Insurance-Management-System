// Function to toggle the sidebar (shrink or expand)
function toggleSidebar() {
  const sidebar = document.getElementById('sidebar');
  const mainContent = document.querySelector('.main-content');

  // Toggle the 'collapsed' class on the sidebar and 'stretched' class on the main content
  sidebar.classList.toggle('collapsed');
  mainContent.classList.toggle('stretched');

  // Save the new state in localStorage
  const isShrunk = sidebar.classList.contains('collapsed');
  localStorage.setItem("sidebarShrunk", isShrunk);
}

// Function to clear session storage for admin
function destroyAdminSession() {
  sessionStorage.removeItem('adminUsername');
  sessionStorage.removeItem('adminEmail');
  sessionStorage.removeItem('adminId');
  sessionStorage.removeItem('adminRole'); // Optional: Clear admin-specific data
}

// Logout functionality for admin
const adminLogoutLink = document.getElementById('adminLogoutLink');
if (adminLogoutLink) {
  adminLogoutLink.addEventListener('click', function (event) {
    event.preventDefault();
    destroyAdminSession();  // Clear admin session data
    window.location.href = '../login.html';  // Redirect to login page
  });
}

// // Ensure admin pages are protected, and user is redirected if not logged in
// if (!sessionStorage.getItem('adminUsername')) {
//   window.location.href = '../login.html';  // Redirect to login if admin is not logged in
// }

// Prevent back navigation for admin users, just like regular users
window.onpopstate = function () {
  destroyAdminSession();  // Clear admin session data
  window.location.href = '../login.html';  // Redirect to login
};

// Ensure page state when restored from cache (e.g., user goes back after logging out)
window.onpageshow = function(event) {
  if (!sessionStorage.getItem('adminUsername') && event.persisted) {
    window.location.href = '../login.html';  // Redirect if admin session doesn't exist
  }
};

// Check the sidebar state on page load
document.addEventListener("DOMContentLoaded", function() {
  const sidebar = document.getElementById('sidebar');
  const mainContent = document.querySelector('.main-content');
  
  // Get the sidebar state from localStorage (collapsed or not)
  const isShrunk = localStorage.getItem("sidebarShrunk") === "true";

  // If the sidebar is shrunk, collapse it and stretch the content
  if (isShrunk) {
    sidebar.classList.add('collapsed');
    mainContent.classList.add('stretched'); // Ensure content is stretched when sidebar is collapsed
  } else {
    // Otherwise, make sure the sidebar is expanded and the content is not stretched
    sidebar.classList.remove('collapsed');
    mainContent.classList.remove('stretched');
  }

  // Optionally: Add a listener for hamburger button if there is one
  const hamburgerButton = document.getElementById('hamburgerBtn');
  if (hamburgerButton) {
    hamburgerButton.addEventListener('click', toggleSidebar); // Toggle sidebar when hamburger button is clicked
  }

  // Handle sidebar links: clicking a link will NOT shrink the sidebar unless explicitly toggled
  const sidebarLinks = document.querySelectorAll('.sidebar a');
  sidebarLinks.forEach(link => {
    link.addEventListener('click', function(event) {
      // Prevent the sidebar from automatically shrinking when clicking on links
      if (!sidebar.classList.contains('collapsed')) {
        // Don't toggle if the sidebar is already in the expanded state
        return;
      }
    });
  });
});

// Ensure that the sidebar state is saved to localStorage when toggled
document.querySelector('#hamburgerBtn').addEventListener('click', toggleSidebar);

// Function to toggle the visibility of the dropdown on click
function toggleSubNav() {
  const dropdown = document.querySelector('.user-nav .dropdown-content');
  dropdown.classList.toggle('active'); // Toggle the 'active' class to show/hide the dropdown
}

// Attach the event listener to the user navigation link
document.getElementById('userNav').addEventListener('click', toggleSubNav);



// JavaScript for toggling the dropdown visibility when clicked
function toggleSubNav() {
  const userNav = document.querySelector('.user-nav');
  userNav.classList.toggle('active'); // Toggle active state for showing the dropdown
}

// Attach the click event to the user navigation element
document.getElementById('userNav').addEventListener('click', toggleSubNav);



// Sidebar Toggler Logic
const sidebarToggler = document.querySelector('.sidebar-toggler');
const sidebar = document.querySelector('.sidebar');
const mainContent = document.querySelector('.main-content');

sidebarToggler.addEventListener('click', () => {
  sidebar.classList.toggle('collapsed');
  mainContent.classList.toggle('stretched');
});

// Modal Handling (ID Proof View)
const openDialogButtons = document.querySelectorAll('.view-btn');
const dialogOverlay = document.querySelector('.dialog-overlay');
const closeDialogButton = document.querySelector('.dialog .close-btn');

openDialogButtons.forEach(button => {
  button.addEventListener('click', () => {
    dialogOverlay.style.display = 'flex';
  });
});

closeDialogButton.addEventListener('click', () => {
  dialogOverlay.style.display = 'none';
});

// Sidebar Dropdown Toggle Logic
const dropdownButton = document.querySelector('.user-nav');
const dropdownContent = document.querySelector('.dropdown-content');

dropdownButton.addEventListener('mouseenter', () => {
  dropdownContent.style.display = 'block';
});

dropdownButton.addEventListener('mouseleave', () => {
  dropdownContent.style.display = 'none';
});

// Table Sorting Functionality
function sortTable(columnIndex, tableId) {
  const table = document.querySelector(`#${tableId}`);
  const rows = Array.from(table.rows).slice(1);
  const isAscending = table.getAttribute('data-sort-order') === 'asc';

  rows.sort((rowA, rowB) => {
    const cellA = rowA.cells[columnIndex].innerText;
    const cellB = rowB.cells[columnIndex].innerText;
    const a = isNaN(cellA) ? cellA : parseFloat(cellA);
    const b = isNaN(cellB) ? cellB : parseFloat(cellB);
    return isAscending ? a > b : a < b;
  });

  rows.forEach(row => table.appendChild(row));
  table.setAttribute('data-sort-order', isAscending ? 'desc' : 'asc');
}

// Attach Sorting to Table Headers
document.querySelectorAll('#claimed th').forEach((header, index) => {
  header.addEventListener('click', () => sortTable(index, 'claimed'));
});

// Toggle Section Visibility Logic
document.querySelectorAll('.section').forEach(section => {
  section.classList.remove('active');
});

function showSection(sectionId) {
  const section = document.getElementById(sectionId);
  if (section) {
    section.classList.add('active');
  }
}

document.querySelectorAll('.section-toggle').forEach(button => {
  button.addEventListener('click', () => {
    const sectionId = button.getAttribute('data-section');
    showSection(sectionId);
  });
});

// Toggle Expanded State of User Navigation
document.querySelector('.user-nav').addEventListener('click', function() {
  this.classList.toggle('expanded');
});

// Set Active Tab in Navbar Based on the Current Page
document.addEventListener("DOMContentLoaded", function() {
  const currentPath = window.location.pathname;
  const addUserTab = document.getElementById("addUserTab");
  const deleteUserTab = document.getElementById("deleteUserTab");
  const modifyUserTab = document.getElementById("modifyUserTab");

  [addUserTab, deleteUserTab, modifyUserTab].forEach(tab => tab.classList.remove("active-tab"));

  if (currentPath.includes("add_user")) {
    addUserTab.classList.add("active-tab");
    document.getElementById("tabText").innerText = "Add User";
  } else if (currentPath.includes("delete_user")) {
    deleteUserTab.classList.add("active-tab");
    document.getElementById("tabText").innerText = "Delete User";
  } else if (currentPath.includes("modify_user")) {
    modifyUserTab.classList.add("active-tab");
    document.getElementById("tabText").innerText = "Modify User";
  }

  document.querySelectorAll('.dropdown-content a').forEach(link => {
    link.addEventListener('click', function () {
      document.querySelectorAll('.dropdown-content a').forEach(item => item.classList.remove('active'));
      this.classList.add('active');
    });
  });
});

// Handle Active Class and Opacity Logic for Dropdown Items
const dropdownItems = document.querySelectorAll('.dropdown-content a');

dropdownItems.forEach(item => {
  item.addEventListener('click', function() {
    dropdownItems.forEach(link => link.classList.remove('expanded'));
    this.classList.add('expanded');
  });

  item.addEventListener('mouseenter', function() {
    this.style.opacity = '0.5';
  });

  item.addEventListener('mouseleave', function() {
    if (!this.classList.contains('expanded')) {
      this.style.opacity = '1';
    }
  });
});
