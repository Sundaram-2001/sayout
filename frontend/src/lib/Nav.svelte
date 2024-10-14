<!-- src/lib/Nav.svelte -->
<nav class="bg-gray-800">
  <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
    <div class="flex items-center justify-between h-16">
      <div class="flex-shrink-0 text-white text-xl">My App</div>
      <div class="hidden sm:block ml-auto">
        <div class="flex space-x-4">
          <a
            href="/home"
            class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >Home</a
          >
          <a
            href="/aboutus"
            class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >About Us</a
          >
          <a
            href="/contact"
            class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >Contact</a
          >
          <a
            href="/register"
            class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >Register</a
          >
          <a
            href="/login"
            class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
            >Login</a
          >
        </div>
      </div>
    </div>
  </div>
</nav>

<slot></slot>
<!-- This allows page content to be inserted here -->
