'use client'
import dynamic from "next/dynamic";
import DropdownList from '../TourTypes'; // Adjust the path accordingly
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import {GetOne} from "@/data/tours";






const EditTourContent = ({ params }) => {
  const router = useRouter();
  console.log('params:', params);
  const tour = GetOne(params.id);



  if (!tour) return null;
  console.log("tour", tour)

  console.log('tourId:', tourId);
  const options = [
    'Adventure',
    'City',
    'Cultural',
    'Museum',
    'Wildlife',
    'Beach',
    'Relax',
    'Ski',
    'Sport',
    'Trekking',
    'Wellness',
    'Cooking',
  ];

    const baseURL = process.env.NEXT_PUBLIC_API_URL;
    const [formData, setFormData] = useState({    
      tag: '',
      title: '',
      location: '',
      price: '',
      latitude: '',
      longitude: '',
      minimum_duration: '',
      group_size: '',
      number_of_reviews: '',
      reviews_comment: '',
      overview: '',
      cancellation_policy: '',
      important_information: '',
      additional_information: '',
      whats_included: '',
      highlights: '',
    });

    // useEffect(() => {
    //   // Fetch user data from the API when the component mounts
    //   const fetchTour = async () => {
    //     try {
          
    //       const response = await fetch(baseURL + `/tours/7`);
    //       if (!response.ok) {
    //         throw new Error('Failed to fetch tour data');
    //       }
    //       const tourData = await response.json();
          
    //       setFormData(tourData.data);
    //     } catch (error) {
    //       console.error('Error fetching tour data:', error.message);
    //     }
    //   };
  
    //   fetchTour();
    // }, [tourId]);
    const handleChange = (e) => {
      setFormData({ ...formData,
        [e.target.name]: e.target.value
      });
    };

    if (!formData) {
      return 

    }

  // handle change and submit Section
  const [selectedOption, setSelectedOption] = useState('');
  const [redirect, setRedirect] = useState(false);
  const { push } = useRouter();

  const submit = async (e) => {
    e.preventDefault();

    tourId = 7;
    const res = await fetch(baseURL + '/tours/' + tourId, {
      method: "PUT",
      credentials: 'include',
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    });

    const body = await res.json();
    console.log("data:", body)
    setRedirect(true);
  };


  useEffect(() => {
    // This effect will run when redirect state changes
    if (redirect) {
      // Perform the redirection
      push('/tour-list-v2');
    }
  }, [redirect, push]);

  const handleSelect = (value) => {
    setSelectedOption(value);
    // You can perform any additional actions based on the selected option here
  };
  return (
    <form className="row y-gap-20" onSubmit={submit}>
      <div className="row x-gap-20 y-gap-20">
        <div className="col-12">
        <div className="form-input" style={{backgroundColor: '#cccccc', border: '2px solid #ddd', borderRadius: '4px', width:'200px', height:'40px'}}>
          <DropdownList className="lh-1 text-16 text-light-1" options={options} onSelect={handleSelect} >
            </DropdownList>
          </div>
        </div>
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='tag' required onChange={handleChange} value={formData?.tag}   />
            <label className="lh-1 text-16 text-light-1">Tag: "LIKELY TO SELL OUT*"</label>
          </div>
        </div>
        {/* End Tag */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='title' required onChange={handleChange} value={formData?.title} />
            <label className="lh-1 text-16 text-light-1">Title: "Mountain Hiking"</label>
          </div>
        </div>
        {/* End Title */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='location' required onChange={handleChange} value={formData?.location} />
            <label className="lh-1 text-16 text-light-1">Location: "City Center"</label>
          </div>
        </div>
        {/* End Location */}

        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='price' required onChange={handleChange} value={formData?.price} />
            <label className="lh-1 text-16 text-light-1">Price: 79.99</label>
          </div>
        </div>
        {/* End Price */}


        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='latitude' required onChange={handleChange} value={formData?.latitue} />
            <label className="lh-1 text-16 text-light-1">Latitude: 34.0522</label>
          </div>
        </div>
        {/* End Latitude */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='longitude' required onChange={handleChange} value={formData?.longitude}/>
            <label className="lh-1 text-16 text-light-1">Longitude: 34.0522</label>
          </div>
        </div>
        {/* End Longitude */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='minimum_duration' required onChange={handleChange} value={formData?.minimum_duration} />
            <label className="lh-1 text-16 text-light-1">Minimum Duration in hours: 100</label>
          </div>
        </div>
        {/* End Minimum Duration */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='group_size' required onChange={handleChange} value={formData?.group_size} />
            <label className="lh-1 text-16 text-light-1">Group Size: 10</label>
          </div>
        </div>
        {/* End Group Size */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='number_of_reviews' required onChange={handleChange} value={formData?.number_of_reviews} />
            <label className="lh-1 text-16 text-light-1">Number of Reviews: 30</label>
          </div>
        </div>
        {/* End Number of Review */}
        <div className="col-12">
          <div className="form-input ">
            <input type="text" name='reviews_comment' required onChange={handleChange} value={formData?.reviews_comment} />
            <label className="lh-1 text-16 text-light-1">Reviews Comment: "Exellent!"</label>
          </div>
        </div>
        {/* End Number of Reviews Comment */}

        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5} name='overview' defaultValue={""} onChange={handleChange} value={formData?.overview} />
            <label className="lh-1 text-16 text-light-1">Overviw: long Text</label>
          </div>
        </div>
        {/* End Overview */}
        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5} name='cancellation_policy' defaultValue={""} onChange={handleChange} value={formData?.cancellation_policy} />
            <label className="lh-1 text-16 text-light-1">Cancellation Policy: long Text</label>
          </div>
        </div>
        {/* End Overview */}
        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5} name='whats_included' defaultValue={""} onChange={handleChange} value={formData?.whats_included} />
            <label className="lh-1 text-16 text-light-1">
            What's Included: Use Comma "," separated values

          </label>
          </div>
        </div>
        {/* End Whats Included */}
        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5}  name='highlights' defaultValue={""} onChange={handleChange} value={formData?.highlights}/>
            <label className="lh-1 text-16 text-light-1">
            Highlights: Use Comma "," separated values
              </label>
          </div>
        </div>
        {/* End Whats Included */}
        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5} name='important_information' defaultValue={""} onChange={handleChange} value={formData?.important_information} />
            <label className="lh-1 text-16 text-light-1">
            Important Inforamtion Use Comma "," separated values
            </label>
          </div>
        </div>
        <div className="col-12">
          <div className="form-input ">
            <textarea required rows={5} name='additional_information' defaultValue={""} onChange={handleChange} value={formData?.additional_information} />
            <label className="lh-1 text-16 text-light-1">
            Additonal Information: Use Comma "," separated values
              </label>
          </div>
        </div>
        {/* End Content */}
 
        <div className="d-inline-block pt-30">
          <button className="button h-50 px-24 -dark-1 bg-blue-1 text-white">
            Save Changes <div className="icon-arrow-top-right ml-15" />
          </button>
        </div>
        {/* End Content */}

      </div>
    </form>
  );
}


export default dynamic(() => Promise.resolve(EditTourContent), {
  ssr: false,
});
