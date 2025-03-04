import React, { useState, useEffect } from "react";
import style from "../../style/bookModal.module.css";
import { useMutation } from "@tanstack/react-query";
import { addBook } from "../../react-queries/api/books";

const AddBookModal = ({ setShowModal }) => {
  const [title, setTitle] = useState("");
  const [isbn, setISBN] = useState("");
  const [authors, setAuthor] = useState("");
  const [publisher, setPublisher] = useState("");
  const [version, setVersion] = useState("");
  const [totalCopies, setTotalCopies] = useState(0);
  const [availableCopies, setAvailableCopies] = useState(0);
    const bookMutation =  useMutation ({ 
        mutationFn:addBook,
        mutationKey:"addBook"
    } )
  const [errors, setErrors] = useState({
    title: "",
    isbn: "",
    author: "",
    publisher: "",
    version: "",
    totalCopies: "",
    availableCopies: "",
  });
const validateForm = () =>{
    let formErrors = { ...errors}
  let valid=true

    if ( totalCopies != availableCopies ) {
        formErrors.availableCopies  = "Available Copies Should be Equal to the Total Copies"
        valid = false
    }
    setErrors(formErrors)
    return valid
}

  const handleSubmit = (e) => {
    e.preventDefault();
   
    if (validateForm()) {
      
      console.log("Form is valid. Submit the data here.");
      bookMutation.mutate({title , authors , isbn, version , publisher , totalCopies , availableCopies })
    //   setShowModal(false);
    } else {
      console.log("Form has errors, please fix them.");
    }
  };

  return (
    <div>
      <div className={style.modalOverlay}>
        <div className={style.modal}>
          <h2 className={style.modalTitle}>Add a New Book</h2>
          <form className={style.modalForm} onSubmit={handleSubmit}>
            {/* ISBN Field */}
            <label className={style.label}>ISBN</label>
            <input
              type="text"
              value={isbn}
              className={`${style.inputField} ${errors.isbn ? style.inputError : ""}`}
              placeholder="Enter ISBN"
              onChange={(e) => setISBN(e.target.value)}
              required
            />
            {errors.isbn && <p className={style.errorText}>{errors.isbn}</p>}

            {/* Title Field */}
            <label className={style.label}>Title</label>
            <input
              type="text"
              value={title}
              className={`${style.inputField} ${errors.title ? style.inputError : ""}`}
              placeholder="Enter Title"
              onChange={(e) => setTitle(e.target.value)}
              required
            />
            {errors.title && <p className={style.errorText}>{errors.title}</p>}

            {/* Author Field */}
            <label className={style.label}>Author</label>
            <input
              type="text"
              value={authors}
              className={`${style.inputField} ${errors.author ? style.inputError : ""}`}
              placeholder="Enter Author"
              onChange={(e) => setAuthor(e.target.value)}
              required
            />
            {errors.author && <p className={style.errorText}>{errors.author}</p>}

            {/* Publisher Field */}
            <label className={style.label}>Publisher</label>
            <input
              type="text"
              value={publisher}
              className={`${style.inputField} ${errors.publisher ? style.inputError : ""}`}
              placeholder="Enter Publisher"
              onChange={(e) => setPublisher(e.target.value)}
              required
            />
            {errors.publisher && <p className={style.errorText}>{errors.publisher}</p>}

            {/* Version Field */}
            <label className={style.label}>Version</label>
            <input
              type="text"
              value={version}
              className={`${style.inputField} ${errors.version ? style.inputError : ""}`}
              placeholder="Enter Version"
              onChange={(e) => setVersion(e.target.value)}
              required
            />
            {errors.version && <p className={style.errorText}>{errors.version}</p>}

            {/* Total Copies Field */}
            <label className={style.label}>Total Copies</label>
            <input
              type="number"
              value={totalCopies}
              className={`${style.inputField} ${errors.totalCopies ? style.inputError : ""}`}
              placeholder="Enter Total Copies"
              onChange={(e) => setTotalCopies(parseInt(e.target.value))}
              required
            />
            {errors.totalCopies && <p className={style.errorText}>{errors.totalCopies}</p>}

            {/* Available Copies Field */}
            <label className={style.label}>Available Copies</label>
            <input
              type="number" 
              value={availableCopies}
              className={`${style.inputField} ${errors.availableCopies ? style.inputError : ""}`}
              placeholder="Enter Available Copies"
              onChange={(e) => setAvailableCopies(parseInt(e.target.value))}
              required
            />
            {errors.availableCopies && <p className={style.errorText}>{errors.availableCopies}</p>}

            {/* Action Buttons */}
            <div className={style.modalActions}>
              <button type="submit" className={style.btnPrimary}>Add Book</button>
              <button type="button" onClick={() => setShowModal(false)} className={style.btnSecondary}>Cancel</button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default AddBookModal;
