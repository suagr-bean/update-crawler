
# Blueprint: Audio Player Vue App

## 1. Application Overview

This is a single-page application (SPA) built with Vue.js. Its primary purpose is to allow users to browse lists of audio content, view details about them, and play the audio tracks through a persistent player at the bottom of the screen.

## 2. Project Outline & Features

### 2.1. Core Technologies

*   **Framework**: Vue.js 3
*   **Syntax**: Composition API (`<script setup>`)
*   **Build Tool**: Vite
*   **HTTP Client**: Axios for fetching data from an API.

### 2.2. Component Architecture

*   `App.vue`: The root component. It manages the overall layout, view switching, and state sharing between the detail view and the player.
*   `Top.vue`: A placeholder for a top navigation bar or header.
*   `Bottom.vue`: The main navigation bar at the very bottom, allowing users to switch between the "Home" and "Sub" views.
*   `playbutton.vue`: The persistent audio player. It receives track info and displays the title and the HTML5 `<audio>` controls. It is always visible after a track is selected.
*   `home.vue`: The main view, displays a list of content items. Clicking on an item navigates the user to the `detail.vue`.
*   `sub.vue`: A secondary view, accessible from the bottom navigation.
*   `detail.vue`: The detail view. It fetches and displays a list of playable audio tracks based on the item selected in `home.vue`.

### 2.3. Design & Layout

*   **Main Layout**: A flexbox column layout managed by `App.vue` (`display: flex, flex-direction: column`).
*   **Content Area**: The central part (`<main class="show">`) is scrollable (`overflow-y: auto`) to accommodate long lists of content.
*   **Fixed Footer**: A `<footer>` element is fixed to the bottom of the viewport (`position: fixed`). This footer acts as a container for both the audio player and the bottom navigation, ensuring they are always visible and accessible.
*   **Player Design**: The `playbutton.vue` component is styled with a light grey background. It has a vertical layout (`flex-direction: column`) to display the track title **above** the audio progress bar and controls.
*   **Content Protection**: The main content area has a `padding-bottom` of `150px` to ensure that no content is ever hidden behind the fixed footer area.

### 2.4. Features & Logic

1.  **View Management**: `App.vue` uses a `ref` (`currentView`) and `v-if` directives to dynamically render the `Home`, `Sub`, or `Detail` components.

2.  **Navigation Flow**:
    *   User starts on the `Home` view.
    *   Clicking an item on the `Home` view triggers the `@detail` event, calling the `ToDetail` method in `App.vue`.
    *   `ToDetail` changes `currentView` to "detail" and stores the item's URL.
    *   The `Detail` view is rendered, receiving the URL as the `apiurl` prop.

3.  **Data Fetching**: The `detail.vue` component uses the `onMounted` lifecycle hook to call the `GetDetail` method, which fetches track data from the `/api/showdetail` endpoint using `axios`.

4.  **Audio Playback Logic**:
    *   In `detail.vue`, each track has a "▶" button.
    *   Clicking this button calls the `Deal(item)` method, which emits a custom event named `@control` upwards, passing the full `item` object of the selected track.
    *   `App.vue` listens for the `@control` event. When it fires, the `Send(item)` method is executed.
    *   `Send` stores the received track `item` into a `ref` called `receive`.
    *   This `receive` object is passed down to the `playbutton.vue` component as the `info` prop.
    *   `playbutton.vue` uses the `info` prop to display the title (`info.title`) and set the audio source (`:src="info.auto_link"`).
    *   The `<audio>` tag includes the `autoplay` attribute to start playing immediately when the `src` changes.

## 3. Current Task: Layout Adjustment

This section outlines the plan for the most recently completed user request.

### 3.1. The Request

The user requested to modify the `playbutton.vue` component to show the track title above the audio player controls. After this change was implemented, the user noted that the taller player component was now overlapping and hiding the content at the bottom of the main view.

### 3.2. The Plan & Execution

1.  **Initial Change**: Modify `playbutton.vue`. Add a `<p class="title">{{ info.title }}</p>` element. Change the flexbox layout to `flex-direction: column` to stack the title on top of the audio player.

2.  **Identify Side Effect**: The change in step 1 increased the overall height of the fixed footer container in `App.vue`.

3.  **Diagnose the Issue**: The `padding-bottom: 120px;` on the main content area (`.show`) was no longer sufficient to create enough space for the taller footer, causing the overlap.

4.  **Implement the Fix**: Modify `App.vue`. Increase the bottom padding on the `.show` style rule from `120px` to **`150px`**. This provides the necessary clearance to prevent any content from being obscured by the now-taller footer.

5.  **Confirmation**: The overlap issue is resolved. The layout now correctly accommodates the player with its title.
