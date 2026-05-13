# User Data

The text editor provides an option to add user data to select lines or every line
in a document. This mechanism is useful when the editor is part of an IDE
where line operations like setting breakpoints are required.

User data is an abstract void pointer that must be managed externally. A number of
APIs are available to manage and use user data. First there are the **SetUserData**
and **GetUserData** methods to attach user data to a specified line or retrieve it.
This is useful when only select lines need user data. Once data is attached, it will
stay  with that line until the line is deleted. Inserting or deleting lines has no
effect on the association. Undo will not restore user data.

There is also a mechanism to attach callbacks (**SetInsertor** and **SetDeletor**)
for when lines are inserted or deleted in a document. This can be used to allocate
and deallocate user data on each line. Given that these callbacks are independent, the
Deletor callback could also be used to manage user data on select lines like removing
breakpoints from deleted lines. The callbacks can be turned off by calling (**SetInsertor** and **SetDeletor**) again with nullptr as the argument. If the callbacks are in place when loading text, you'll get a callback for each line. If
they are in place when clearing or disposing an editor instance, you get a callback
for each deleted line. It is therefore important to (re)set callbacks at the right
time to get the desired effect.

User data pointers are also passed to the optional decorator callback allowing
integrators to quickly access their data when rendering line decorations. Last but
not least, there is a user data iterator (**IterateUserData**) which calls a
callback on each line that has user data attached.

So to use these mechanisms to manage user data on each line you can use a
pattern like:

```c++
struct UserData {
	int somethingImportant;
	float somethingEvenMoreImportant;
};

editor.SetInsertor([]([[maybe_unused]] int line) {
	auto data = new UserData();
	return static_cast<void*>(data);
});

editor.SetDeletor([]([[maybe_unused]] int line, void* userData) {
	auto data = static_cast<UserData*>(userData);
	delete data;
});

editor.SetLineDecorator(40.0f, [](TextEditor::Decorator& decorator) {
	auto data = static_cast<UserData*>(decorator.userData);

	if (data.somethingImportant) {
		// do something
	}
});
```

To only use these mechanisms for notifications, use a pattern like this:

```c++
editor.SetInsertor([](int line) {
	std::cout << "Inserted line " << line << std::endl;
	// we must return a value so we just use nullptr
	return nullptr;
});

editor.SetDeletor([](int line, [[maybe_unused]] void* userData) {
	std::cout << "Deleted line " << line << std::endl;
});

To only apply these mechanisms to a few lines, use a pattern like:

```c++
struct BreakPoint {
	bool active = false;
};

auto breakpoint = new BreakPoint();
editor.SetUserData(10, static_cast<void*>(breakpoint));

editor.SetDeletor([]([[maybe_unused]] int line, void* userData) {
	if (userData) {
		auto data = static_cast<BreakPoint*>(userData);
		delete data;
	}
});

editor.SetLineDecorator(40.0f, [](TextEditor::Decorator& decorator) {
	if (decorator.userData) {
		auto data = static_cast<UserData*>(decorator.userData);
		// render breakpoint based on user data
	}
});

// summarize all breakpoints
editor.IterateUserData([](int line, void* userData){
	auto data= static_cast<BreakPoint*>(userData);
	// do something
});
```