'use strict';
// ====================THE SIDEBAR=============================//
/**Here, a class of 'menu-item' is applied to all the anchor tags in the sidebar starting from 
 * HTML LINE NUMBER 54 TO 78 The first 'menu-item' has the class of 'active' at CSS line no. 335. 
 * this 'active' class makes the background of its parent element transparent.
 */

// This code selects all anchor tags with the class of 'menu-item'
const menuItems = document.querySelectorAll('.menu-item');

// This function (handleActivity) removes the 'active' class from the each 'menu-item'
const handleActivity = () => {
    menuItems.forEach(item => {
        item.classList.remove('active');
    });
}

// Here, a 'forEach' is used to loop through all the 'menu-item'
menuItems.forEach(item => {
    /**Here, a click listener is created. So for 'menu-item' that is clicked, a classList of
     * 'active' is added, and the handleActivity(JS Line No: 25) function is called to remove the 'active' class
     * from the previous 'menu-item'
     */
    item.addEventListener('click', function() {
        handleActivity();
        item.classList.add('active');
    });
});


// ============THE EDIT POPUP======================= (The div starts from HTML LINE: 691 TO 725);

/**There are a total of five edit buttons at HTML LINE NO: 274, 341, 413, 476, 555 respectively.
 * the code below selects all of those buttons
*/
const EditButton = document.querySelectorAll('.edit-btn');

/**Here, the popup container with an ID of 'edit-popup' (at HTML LINE NO: 691) is selected */
const editPopup = document.getElementById('edit-popup');

/**The popup container contains a direct child (which contains all the edit-options) with an ID of 'edit-card' 
 * (at HTML LINE NO: 692); The code below selects that card.
 */
const editCard = document.getElementById('edit-card');

// A button with an ID of 'close-edit-popup-btn' is selected at HTML LINE NO: 693. This button closes the popup.
const closeEditCard = document.getElementById('close-edit-popup-btn');

/**Here, i used a 'for' loop to loop through all the 'EditButton' (JS LINE: 36)*/
for(let i = 0; i < EditButton.length; i++) {
    // Storing the buttons in a variable called 'editBtn'
    const editBtn = EditButton[i];

    /**An event listener (click) is created here, which listens for an event whenever any of the editBtn
     * is clicked.
     */
    editBtn.addEventListener('click', () => {
        /**Whenever an 'editBtn' (JS LINE NO: 52) is clicked, a classList of 'edit-popup-is-active'
         * (CSS LINE NO: 799) is added to the document; this displays the popup that contains all the options 
         * for editing a post
         */
        editPopup.classList.add('edit-popup-is-active');

        // Clip-path is also added to the edit-card (CSS LINE NO: 820)
        editCard.classList.add('edit-card-clip-path-is-active');
    })
}

/** Remember the button we selected? the button that closes the popup-container (JS LINE NO: 47)
 *  Here, a function called 'removeEditCard()' that removes the edit-container is created, and then assigned to 
 *  the closeEditCard (JS LINE NO: 85);
 */
const removeEditCard = () => {
    editCard.classList.remove('edit-card-clip-path-is-active');
    editPopup.classList.remove('edit-popup-is-active');
}

/**I created a fucntion that checks if a user is clicking outside an area that does not contain the edit
 * options, but contains the class of 'edit-popup-container' at HTML LINE NO: 691. So basically, what this
 * fucntion does is remove the popup enitrely once it detects that a user is clicking outside the card that
 * contains all the edit options
 */
const closeEditPopup = (event) => {
    if(event.target.classList.contains('edit-popup-container')) {
        editPopup.classList.remove('edit-popup-is-active');
    }
}

// The function above is then added to the popup div (at HTML LINE NO: 691).
editPopup.addEventListener('click', closeEditPopup);

// The removeEditCard function is added to the button that closes the popup-container
closeEditCard.addEventListener('click', removeEditCard);



//=====================ADD AND REMOVE FROM BOOKMARK=====================//
// (HTML LINE NO: 728 TO 736);

/**1. there are also a total of five bookmark buttons (at HTML LINE NO: 308, 383, 443, 521, 618 RESPECTIVELY.)
 *    the code below (JS LINE NO: 103) selects all of those buttons with a class of 'bookmark-btn' using a
 *    queryselectorall   
 */
const bookmarkBtn = document.querySelectorAll('.bookmark-btn');

// The button that closes the bookmark (at HTML LINE NO: 734) is selected below
const closeBookmarkBtn = document.getElementById('close-bookmark');

// The bookmark popup-container is selected by its ID 'bookmark-popup' (at HMTL LINE NO: 728)
const bookmarkPopup = document.getElementById('bookmark-popup');

/**Inside of the bookmar-container, a paragraph with an ID of 'JS-editable' is selected below.
 * The paragraph is a flexible type which changes its inner-content depending on the action 
 * of the user on the bookmark-button.
 */
const messageChange = document.getElementById('JS-editable');

/** The bookmarkBtn (at JS LINE NO: 104) is looped through uisng a 'forEach' loop. For each iteration, a
 * function (with an arguement called 'bookmarkButton') is passed.
*/
bookmarkBtn.forEach(bookmarkButton => {
    // for each button, an event listener is added
    bookmarkButton.addEventListener('click', () => {
        /**Here, whenever i click a bookmark button, i want to do two things:
         * 1. I want to add a classList called 'bookmark-btn-is-checked'(at CSS LINE NO: 955) to the button.
         * 2. Then i want to toggle the classList added above 
        */
        bookmarkButton.classList.toggle('bookmark-btn-is-checked');

        /**I created a condition here that tests if a bookmarkButton contains a classList of bookmark-btn-is-checked.
         */
        if (bookmarkButton.classList.contains('bookmark-btn-is-checked')) {
            /**If it contains, 
             * 1. it should add a classList of 'bookmark-container-is-active' (at CSS LINE NO: 1155) to the 
            bookmark-popup that wAs earlier selected (at JS LINE NO: 110)*/
            bookmarkPopup.classList.add('bookmark-container-is-active');
            // 2. The message should update the changes that has been made by the user, and then displays it to the user.
            messageChange.innerHTML = `Gist has been added successfully to your bookmarks.`

            /**But if it contains the default bookmarkButton (at CSS LINE NO: 951) */
        } else if (bookmarkButton.classList.contains('bookmark-btn')) {
            // 1. I still want to add the classList of 'bookmark-container-is-active' to the popup
            bookmarkPopup.classList.add('bookmark-container-is-active');

            // 2. But, i want to update the changes that has been made, and then display it as a message to the user.
            messageChange.innerHTML = `Gist has been removed from your bookmarks`;        
        }
        
        // Here, the closeBookmarkBtn is assigned a istener, and a function that closes the popup, whenever it's clicked
        closeBookmarkBtn.addEventListener('click', () => {
            bookmarkPopup.classList.remove('bookmark-container-is-active');
        });
        // Set a timeout function that determines how long the popup stays on the page
        setTimeout(() => {
            bookmarkPopup.classList.remove('bookmark-container-is-active');
        }, 3000); 
    })
})


// =========================CUSTOMIZE USER EXPERIENCE======================= 
// Start: HTML LINE NO: 637 and END: HTML LINE NO: 688

/** Here, i gave an ID of 'theme' to the fifth 'menu-item' (at HTML LINE NO: 74 and 945); 
 * The line of below selects all the menu-item with an ID of 'theme'
*/
const theme = document.querySelectorAll('#theme');

// The theme modal with a class of 'customize-theme' (at HTML LINE NO: 637) is also selected 
const themeModal = document.querySelector('.customize-theme');

// Just like the function at JS LINE NO: 78, the modal closes once a user clicks outside of the theme-card
const closeModal = (e) => {
    if(e.target.classList.contains('customize-theme')) {
        themeModal.classList.remove('theme-card-is-active')
    }
}

const closeThemeModal = document.getElementById('close-theme-modal');
closeThemeModal.addEventListener('click', closeModal)

// The function above (closeModal()) is then called and assigned to the themeModal
themeModal.addEventListener('click', () => {
    themeModal.classList.remove('theme-card-is-active')
});

/**Using a forEach loop, for every menu-item (with an ID of 'theme') that is clicked, a function with the
 * argument called 'theme', is passed */
theme.forEach(theme => {
    theme.addEventListener('click', () => {
        /** Whenever a user clicks a theme button, we want to add a classList of 'theme-card-is-active'
         * at (CSS LINE NO: 991) which allows the popup action of the customize-theme-card
         */
        themeModal.classList.add('theme-card-is-active');
    });
})

/**Here. the different font-sizes (as seen in HTML LINE NO: 647 TO 651) are selected using a queryselectorall, and
 * then stored in a variable
*/
const fontSizes = document.querySelectorAll('.choose-size span');

/**The line of code below selects the 'root' pseudo-class (at CSS LINE NO: 15). In this root class, different
 * variables have been stored (including colors, padding, border-radius e.t.c) 
 */
let root = document.querySelector(':root');

/**This function removes the 'active' class (at CSS LINE NO: 1035) from each of the font-size selectors.
 * the 'active' class changes the background of the font-size selector to a defined color.
*/
const removeSelector = () => {
    fontSizes.forEach(size => {
        size.classList.remove('active');
    });
}

/**A forEach loop is used to loop through all the fontSize selectors, an event listener is added 
 * to each one, and then a function (with an arguement called 'size') is assigned to each selector */
fontSizes.forEach(size => {
    size.addEventListener('click', () => {
        /**whenever a user clicks one of the selectors, 
         * 1. I want to remove the class of 'active' from the previous and
         * next selector, but the clicked selector should maintain the classList of 'active'.
         */
        removeSelector();

        //2. Set an empty variable called fontSize using the 'let' keyword, since it's going to be flexible
        let fontSize;

        /**3. I want to toggle the 'active' class on each of the selectors */
        size.classList.toggle('active');

        // if the size selector contains a classList of 'font-size-1' (at HTML LINE NO: 647)
        if(size.classList.contains('font-size-1')) {
            // Then i want to set the fontSize (defined above: JS LINE NO: 223) to be 10px
        fontSize = '10px';
        /**after i set the fontSize to the defined size, I want to adjust the position of sidebar
         * so that it matches the fontSize of the document.
        */
        root.style.setProperty('----sticky-top-left', '5.4rem'); //at CSS LINE NO: 45
        root.style.setProperty('----sticky-top-right', '5.4rem'); //at CSS LINE NO: 46
        } else if (size.classList.contains('font-size-2')) {
            fontSize = '13px';
            root.style.setProperty('----sticky-top-left', '5.4rem');
            root.style.setProperty('----sticky-top-right', '-7rem');
        } else if (size.classList.contains('font-size-3')) {
            fontSize = '16px'
            root.style.setProperty('----sticky-top-left', '-2rem');
            root.style.setProperty('----sticky-top-right', '-17rem');
        } else if (size.classList.contains('font-size-4')) {
            fontSize = '19px'
            root.style.setProperty('----sticky-top-left', '-5rem');
            root.style.setProperty('----sticky-top-right', '-25rem');
        } else if (size.classList.contains('font-size-5')) {
            fontSize = '22px';
            root.style.setProperty('----sticky-top-left', '-12rem');
            root.style.setProperty('----sticky-top-right', '-33rem');
        }

        // Whichever fontSize we select, we want to set the fontSize of the html document to be = fontSize (JS LINE: 223)
        document.querySelector('html').style.fontSize = fontSize;
    });
});


// SAME LOGIC GOES FOR THE COLOR VARIATION
const colorVariants = document.querySelectorAll('.choose-color span');

const removeActiveClass = () => {
    colorVariants.forEach(colorPicker => {
        colorPicker.classList.remove('active'); //ACTIVE CLASS CSS LINE NO: 1079
    });
}
colorVariants.forEach(color => {
    color.addEventListener('click', () => {
        // FOR EVERY COLOR VARIANT THAT IS CLICKED, A NEW VARIABLE CALLED 'primaryHue' IS CREATED.
        let primaryHue;

        // The 'active' class, which adds a white border to each color variant, is removed from the previous and next
        // variant; but the clicked color maintains the 'active' class.
        removeActiveClass();

        // If the clicked color variant has the class of 'color-1', we change the hue value
        if (color.classList.contains('color-1')) {
            primaryHue = 252;
        } else if (color.classList.contains('color-2')) {
            primaryHue = 52;
        } else if (color.classList.contains('color-3')) {
            primaryHue = 352;
        } else if (color.classList.contains('color-4')) {
            primaryHue = 152;
        } else if (color.classList.contains('color-5')) {
            primaryHue = 202;
        }

        // This adds the class list of 'active' (at CSS LINE NO: 1079)
        color.classList.add('active');

        /**So whichever variant a user click, the 'root' class is accessed, and then the variable called 
         * '--primary-color-hue' is set to the 'primaryHue' variable
        */
        root.style.setProperty('--primary-color-hue', primaryHue);
    });
});

// Here, the different background variations are selected by their ID
const bg1 = document.getElementById('bg-1');
const bg2 = document.getElementById('bg-2');
const bg3 = document.getElementById('bg-3');

// These empty variables are flexible, and they serve to alter the text-color when the background is altered
let lightColorLightness;
let darkColorLightness;
let whiteColorLightness;

/**Just as the name implies, the handleBgChange() function changes the property of the 'root' class
 * it replaces the white, light, and dark colors with the above variables
 */
const handleBgChange = () => {
    root.style.setProperty('--light-color-lightness', lightColorLightness);
    root.style.setProperty('--white-color-lightness', whiteColorLightness);
    root.style.setProperty('--dark-color-lightness', darkColorLightness);
}


// An event listener is added to the background class. when the first background variable is clicked,
bg1.addEventListener('click', () => {
    // a class of 'active' (at CSS LINE NO: 1105 ---> THIS ADDS A BORDER TO THE CLICKED BG) is added
    bg1.classList.add('active');

    // But then, the same 'active' class is removed from the rest of the background variables
    bg2.classList.remove('active');
    bg3.classList.remove('active');

    // removes the customized changes from the localStorage
    window.location.reload();
});

// same logic as JS LINE NO: 321 IS ALSO APPLIED HERE 
bg2.addEventListener('click', () => {
    // But the lightness are altered. this cuases the changes in text color
    darkColorLightness = '95%';
    whiteColorLightness = '20%';
    lightColorLightness = '15%';

    bg2.classList.add('active');

    bg1.classList.remove('active');
    bg3.classList.remove('active');

    // Then the function handleBgChange() function is called. this causes the replacement of the root variables
    handleBgChange();
});
bg3.addEventListener('click', () => {
    darkColorLightness = '95%';
    whiteColorLightness = '10%';
    lightColorLightness = '0%';

    bg3.classList.add('active');

    bg1.classList.remove('active');
    bg2.classList.remove('active');
    handleBgChange();
});

// The iconbars contain anchor tags with a class of 'menu-item'. The code below selects all of the 'menu-item' class
const fixedIcons = document.querySelectorAll('.iconbar .menu-item');

// This function removes the active class (at CSS LINE NO: 1777) from the 'menu-item'
const handleMenuActivity = () => {
    fixedIcons.forEach(item => {
        item.classList.remove('active')
    });
}

// We loop through all the 'menu-item', and then pass a function
fixedIcons.forEach(item => {
    item.addEventListener('click', function() {
        // for every clicked item, we handleMenuActivity() (i.e. remove the active class)
        handleMenuActivity();
        // and then, we add the active class to the clicked item
        item.classList.add('active');
    });
});


// =========================CREATING A POST================================//
const postHandler = document.getElementById('post-type'); //Create-post-button (at HTML LINE NO: 28)
const postContainer = document.getElementById('create-post-container'); //at HTML LINE NO: 739

// We write a function that adds the classList of 'create-post-container-is-active' (at CSS LINE NO: 1198) to the post-container
const handlePostContainer = () => {
    // THIS ALLOWS THE DISPLAY OF POST-CREATION CONTAINER
    postContainer.classList.add('create-post-container-is-active');
}

// Set the function that closes the post-container whenever a user clicks outside the container
const removePostContainer = (e) => {
    if(e.target.classList.contains('create-post-container'))/**at HTML LINE NO: 739 */ {
        postContainer.classList.remove('create-post-container-is-active');
    }
}

// Add the function to the postHandler
postHandler.addEventListener('click', handlePostContainer);
postContainer.addEventListener('click', removePostContainer);


/**Here, container with and ID of 'dropdown-filter-buttons' is selected (at HTML LINE NO: 764). This dropdown contains all the filter
 * buttons i.e. single-post-button, series-post-button, novel-post.....
 */
const dropdownFilter = document.getElementById('dropdown-filter-buttons');

/**Just above the 'dropdown-filter-buttons', there is a default select button with an ID of 'default' (at HTML LINE NO: 759) */
const postSelectButton = document.getElementById('default');

/**Whenever i click the postSelectButton,  */
postSelectButton.addEventListener('click', function() {
    // I want to toggle the classList of 'filter-buttons-are-active' (at CSS LINE NO: 1321). Thi class displays the dropdownFilterButtons
    dropdownFilter.classList.toggle('filter-buttons-are-active');
})

// Set a function that hides the dropdown
const hideDropdownFilter = () => {
    dropdownFilter.classList.remove('filter-buttons-are-active');
}

/**The code below selects all the filter buttons in the filterButton container (at HTML LINE NO: 764) */
const AllFilterButtons = document.querySelectorAll('.post-type-container .filter-buttons label');

/** We select a p tag with an ID of 'warning'. The purpose of this warningMessage is to tell the user the consequence of 
 * choosing a particular post-type. It is displayed directly under the default button (at HTML LINE NO: 763)
*/
const warningMessage = document.getElementById('warning');

// Select all the post-types (HTML LINE NO: 779, 799, 817, 822, 842, 870, 893)
const PostTypes = document.querySelectorAll('post-types post');

// Select the image and video icons (at HTML LINE NO: 923, 928)
const imgIcon = document.getElementById('imgIcon');
const vidIcon = document.getElementById('vidIcon');

/**Here, we set a function that removes the currentPostType whenever a user clicks on a post-type that is different from the current 
 * post-type. It removes the classList of 'post-active' (atCSS LINE NO: 1733)
 */
const removeCurrentPostType = () => {
    PostTypes.forEach(post => {
        post.classList.remove('post-active');
    })
}

// Select the card that contains all the contents for post creation (at HTML LINE NO: 740)
const creatPostCard = document.getElementById('create-post-card');

// THIS FUNCTION IS EXCLUSIVELY FOR THE NOVEL POST, ONCE A USER SELECTS THE NOVEL POST, THE CARD THEN TAKES THE FULL HEIGHT OF THE PAGE
const handleCardHeight = () => {
    creatPostCard.classList.remove('full-height');
    creatPostCard.classList.remove('mid-height');
    creatPostCard.classList.remove('no-border-radius');
}

// Set a function that handles the visibility of all the other post-types whenever a user clicks on a post-type
const handleVisibility = () => {
    // Here, we select all post-types with data-attribute of 'data-post' (also at HTML LINE NO: 779, 799, 817, 822, 842, 870, 893);
    const post = document.querySelectorAll('[data-post]');
    // Then we loop through all post-types using a forEach loop; and then we pass in a function with an arguement called item.
    post.forEach(item => {
        /**So for every item (basically 'post-type') a user clicks, we want to remove the class of 'post-active' from the other items
         * so that only the clicked item maintains the  'post-active' class.
         */
        item.classList.remove('post-active');
    });
}

/**This function below handles the upload control buttons WE SELECTED EARLIER (at JS LINE NO: 435 AND 436). This function is set for
 * a particular post-type
 */
const handleUploadControlButtons = () => {
    // We set the display of the icons to be 'none'.
    vidIcon.style.display = 'none';
    imgIcon.style.display = 'none';
}

/**Here, we loop through all the filter buttons that we selected earlier (at JS LINE NO: 424) using a forEach loop, and then we pass in
 * a function with an argument called 'button'.
 */
AllFilterButtons.forEach(button => {
    /**For every button (i.e. any filter button) that a user clicks: */
    button.addEventListener('click', () => {
        //1. We want to call back the following functions:

        // A. The hideDropdownFilter(). This function hides the popup that displays all the filter buttons (i.e. single, series ......)
        hideDropdownFilter();
        // B. The handleCardHeight(). This function removes the 'full-height' class (at CSS LINE NO:1737) from all the post-types except the novel-post
        handleCardHeight();
        // C. The handleVisibility(). This function remvoes the 'post-active' from all the post-types except the current post-type
        handleVisibility();
        /**D. The handleUploadControlButtons(). This function sets the display of the image and video icons to none; but it only applies to
         * certain post-types.
        */
        handleUploadControlButtons();

        /**Here, we create some conditions that tests if the ID of the button that a user clicks is equal to that of the post-type */
        // =================SINGLE=========================
        if (button.id === 'Single') {
            // We call the handleUploadControlButtons() since the single post only allows text. No images! No videos!!
            handleUploadControlButtons;
            const singlePost = document.getElementById('single_post').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            // If the IDs match, then we set the update the warning message and display it to the user.
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> Note: In a ${button.id} post, you are only allowed a maximum of 250 words`;
            // =======================SERIES================
        } else if (button.id === 'Series') {
            handleUploadControlButtons;
            const seriesPost = document.getElementById('series_post').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> This post-type is similar to a single post, but in episodes, with a maximum of 50 words per episode.`;
            // ==================NOVEL=====================
        } else if (button.id === 'Novel') {
            handleUploadControlButtons;
            const novelPost = document.getElementById('novel_post').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> The novel post-type is more like a chaptered story, each chapter having a maximum of 1500 word`;
            creatPostCard.classList.add('full-height');
            creatPostCard.classList.add('no-border-radius')
            // =================BOX-OFFICE===================
        } else if (button.id === 'Box-office') {
            handleUploadControlButtons;
            vidIcon.style.display = 'flex';
            const boxOffice = document.getElementById('video_only').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> The box-office post-type is movie-like video-only gist. However, you are allowed to upload one video at a time.`;
            // ===================SOAP-OPERA====================
        } else if (button.id === 'Soap-opera') {
            handleUploadControlButtons;
            const SoapOpera = document.getElementById('video-series').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> Similar to the box-office and the series post-type, you can post multiple videos at a time.`;
            // ==================RADIO=====================
        } else if (button.id === 'Radio') {
            handleUploadControlButtons;
            const Radio = document.getElementById('audio_post').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> This post-type is a live (recordable) audio broadcast`;
            // ==================LIVESTREAM====================
        } else if (button.id === 'LiveStream') {
            handleUploadControlButtons;
            const LiveStream = document.getElementById('live_stream').classList.add('post-active');
            postSelectButton.innerHTML = `${button.id} <i class="fa fa-chevron-right"></i>`;
            warningMessage.innerHTML = `<i class="fa fa-circle"></i> This post-type is a live (recordable) video broadcast`;
            creatPostCard.classList.add('full-height');
        } 
    })
})


/**tHIS FUNCTION IS CREATED EXCLUSIVELY FOR THE SERIES POST. It creates a new text-box whenever we click a button (selected at JS LINE NO:
 * 558) */
const handleSectionCreate = () => {
    // WE SET THE TYPE OF ELEMENT WE WANT TO CREATE. IN THIS CASE, A 'TEXTAREA'
    const textArea = document.createElement('textarea');
    // WE SET A 'name' ATTRIBUTE TO THE TEXTAREA. I.E. name="series"
    textArea.setAttribute('name', 'series');
    // WE ALSO SET THE ROW ATTRIBUTE I.E. THE NUMBER OF ROWS (rows="7")
    textArea.setAttribute('rows', 7);

    // Then we select a container that we'll append the created element. I this case, we select the 'text-container' div (at HTML LINE NO: 784)
    const textContainer = document.getElementById('text-container');
    // APPEND THE CREATED ELEMENT TO THE 'text-container' AS A CHILD ELEMENT
    textContainer.appendChild(textArea);
}

/**SINCE WE NEED TO ASSIGN THE FUNCTION handleSectionCreate(), WE SELECT THE 'add-section-button' (at HTML LINE NO: 809) */
const createSection = document.getElementById('add-section-button');
// The we assign the function
createSection.addEventListener('click', handleSectionCreate);

/**We set a function that loads and displays a file for preview, whenever we upload one. The function comes with an arguement called
 * 'event'
 */
const loadFile = (event) => {
    // We get the file-container by its ID 'output' (at HTML LINE NO: 834)
    let video = document.getElementById('output');
    // We create a URL for video source(src). Here, we target the files at the first position
    video.src = URL.createObjectURL(event.target.files[0]);
    // Then we set a 'control' attribute to the video. This allows a user to be able to play, set playback speed, and so on....
    video.setAttribute('controls', true);
    // Whenever a video is loaded and displayed, we want set the display of the vidIcon to be 'none'
    vidIcon.style.display = 'none';
}

// This object contains different CSS STYLES accessed at JS LINE NO: 610 TO 612
const style = {
    border: '1px solid var(--color-primary)',
    borderRadius: 'var(--card-border-radius)',
    marginTop: '1rem'
}

/**Similar to the handleSectionCreate() function, we also create a function (exclusively for the SOAP-OPERA). This function creates a new
 * 'Add video' section when a user clicks a button.
*/
const CreateVideoSection = () => {
    // In this case, we create the 'div' element
    const videoSection = document.createElement('div');
    // and set the class attribute
    videoSection.setAttribute('class', 'serie');
    // Then we give it a custom HTML markup
    videoSection.innerHTML = `
        <div class="caption-wrapper">
            <textarea id="caption-for-vid-series" cols="30" rows="2"
            placeholder="Add caption"></textarea>
        </div>
        <div class="video-wrapper">
            <div class="series-vid-container">
                <video id="serie-vid"></video>
            </div>
            <input type="file" name="video" accept="video/*" id="seriesVids" hidden>
            <label for="" style="display: flex; align-items: center; gap: .35rem; 
            padding: 0 var(--card-padding) .5rem">
                <i class="fas fa-video"></i>
                <h1 style="font-size: 16px;">Upload video</h1>
            </label>
        </div>
    `;
    // remember the 'style' object we created earlier? We use it to style the videoSection. So basically, everytime we create a new section, 
    // we add these styles to it.
    videoSection.style.border = style.border;
    videoSection.style.borderRadius = style.borderRadius;
    videoSection.style.marginTop = style.marginTop;

    // Then we select the container that we'll append the created element (at HTML LINE NO: 843)
    const seriesWrapper = document.getElementById('series-wrapper');
    // Append the created element to the 'series-wrapper' as a child element
    seriesWrapper.appendChild(videoSection);
}

// Select the button that will carry out the 'CreateVideoSection()' function
const addVideoButton = document.getElementById('add-vid-button');
// The assign the function to the button
addVideoButton.addEventListener('click', CreateVideoSection);


// Generate a Unique ID
// let generateUniqueID = () => {
//     let StringID = () => {
//         return Math.floor((1 + Math.random()) * 0x10000)
//         .toString(16).substring(1);
//     }
//     return StringID() + StringID() + '-' + StringID() + '-' + StringID() + '-' + StringID() + '-' + StringID() + StringID() + StringID();
// }


// ====================Code for lIVESTREAM POST======================
let videoElement = null;
let startTriggerButton = document.getElementById('start-trigger');

let startVideo = () => {
    let camAvailable = navigator.mediaDevices && navigator.mediaDevices.getUserMedia;
        if(camAvailable) {
            videoElement = document.getElementById('video');
            navigator.mediaDevices.getUserMedia({video: true, audio: true}).then(function(stream){
                videoElement.srcObject = stream;
                videoElement.play()
        });
    };
};

let pauseVideo = () => {
    videoElement.pause();
}

let audioElement = null;

let startAudio = () => {
    let audAvailable = navigator.mediaDevices && navigator.mediaDevices.getUserMedia;
        if(audAvailable) {
            audioElement = document.getElementById('audio');
            navigator.mediaDevices.getUserMedia({audio: true}).then(function(stream){
                audioElement.srcObject = stream;
                audioElement.play();
        });
    };
};

let pauseAudio = () => {
    audioElement.pause();
}